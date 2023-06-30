package logger

import (
	"context"
<<<<<<< HEAD
	"hslj/config"
	"hslj/pkg/util"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"fmt"
	"hslj/config"
	"hslj/pkg/util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type (
	Level int
)

const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var levelName = []string{
	LevelFatal:   "output",
	LevelDebug:   "debug",
	LevelInfo:    "info",
	LevelWarning: "warn",
	LevelError:   "error",
}

var (
	logFileDir  = util.GetStringValue(config.Vipper.Get("Logger.LogFileDir"))
	projectName = util.GetStringValue(config.Vipper.Get("Logger.ProjectName"))
)

func getProjectLogPath(projectName string, logFileName string) string {
	if projectName == "" {
		panic("please check your config, not set Logger.ProjectName")
	}
	projectLogPath := filepath.Join(logFileDir, projectName)
	if logFileName == "" {
		return projectLogPath
	}
	return filepath.Join(projectLogPath, logFileName)
}

const (
	TraceIdKey = "trace_id"
)

// 获取trace_id
func GetTraceId(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(TraceIdKey)
}

// 生成trace_id
func GenerateTraceId() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 生产 ctx
func NewTraceIdContext() context.Context {
	return context.WithValue(context.Background(), TraceIdKey, GenerateTraceId())
}

var _fileLogWriter = newFileWriter()

// fileLogWriter implements LoggerInterface.
// It writes messages by lines limit, file size limit, or time frequency.
type fileLogWriter struct {
	//sync.RWMutex // write log order by order and  atomic incr maxLinesCurLines and maxSizeCurSize

	logSync    sync.Mutex
	accessSync sync.Mutex

	// The opened file
	_logFile       map[Level]*os.File
	_accesslogFile *os.File

	// logs map cache
	_logs      map[Level]*logger
	_accesslog *logger

	LogFilePath string

	// Rotate daily
	Daily bool `json:"daily"`
	//dailyOpenDate int
	dailyOpenTime time.Time

	accessLogFileTime time.Time
	logFileTime       map[Level]time.Time

	Perm string `json:"perm"`
}

// Logger defines the behavior of a log provider.
type Logger interface {
	Init() error
	Log(level Level) *logger
	Accesslog() *logger
}

// newFileWriter create a FileLogWriter returning as LoggerInterface.
func newFileWriter() Logger {
	w := &fileLogWriter{
		Daily:       true,
		Perm:        "0664",
		_logs:       make(map[Level]*logger, 0),
		_logFile:    make(map[Level]*os.File, 0),
		logFileTime: make(map[Level]time.Time, 0),
		LogFilePath: getProjectLogPath(projectName, ""),
	}
	return w
}

func (w *fileLogWriter) Accesslog() *logger {
	w.accessSync.Lock()
	defer w.accessSync.Unlock()
	if w._accesslog == nil || w.checkAccessLogNeedRotate() {
		err := w.createAccessLogFile()
		if err != nil {
			fmt.Printf("err %#v", err.Error())
		}
	}
	return w._accesslog
}

func (w *fileLogWriter) checkAccessLogNeedRotate() bool {
	return time.Now().Day() != w.accessLogFileTime.Day()
}

func (w *fileLogWriter) Log(level Level) *logger {
	w.logSync.Lock()
	defer w.logSync.Unlock()
	if w._logs[level] == nil || w.checkLogNeedRotate(level) {
		err := w.createLogFile(level)
		if err != nil {
			fmt.Printf("err %#v", err.Error())
		}
	}

	return w._logs[level]
}

func (w *fileLogWriter) checkLogNeedRotate(level Level) bool {
	if time.Now().Hour() != w.logFileTime[level].Hour() {
		return true
	}
	return false
}

func (w *fileLogWriter) Init() error {
	//检查是否存在项目目录，不存在则创建目录
	projectPath := filepath.Join(logFileDir, projectName)
	if !w.isExist(projectPath) {
		err := os.MkdirAll(projectPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func (w *fileLogWriter) isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func (w *fileLogWriter) createLogFile(level Level) error {
	basePath := w.LogFilePath
	t := time.Now()
	year, month, day := t.Date()
	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}
	perm, err := strconv.ParseInt(w.Perm, 8, 64)
	if err != nil {
		return err
	}
	hour := t.Hour()
	fileName := fmt.Sprintf("%s-%04d%02d%02d-%02d-%s.log", projectName, year, month, day, hour, levelName[level])
	fname := filepath.Join(basePath, fileName)
	//正常日志
	logFileFd, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	// Make sure file perm is user set perm cause of `os.OpenFile` will obey umask
	os.Chmod(fname, os.FileMode(perm))

	w.logFileTime[level] = t

	//close the old file
	if w._logFile[level] != nil {
		w._logFile[level].Close()
	}
	w._logFile[level] = logFileFd

	if w._logs[level] != nil {
		w._logs[level]._log.SetOutput(logFileFd)
	} else {
		w._logs[level] = &logger{_log: log.New(logFileFd, "", log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
	}

	return nil
}

func (w *fileLogWriter) createAccessLogFile() error {
	basePath := w.LogFilePath
	t := time.Now()
	year, month, day := t.Date()
	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}
	perm, err := strconv.ParseInt(w.Perm, 8, 64)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%04d%02d%02d-access.log", year, month, day)
	fname := filepath.Join(basePath, fileName)

	accessLogFileFd, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(perm))
	if err != nil {
		return err
	}

	os.Chmod(fname, os.FileMode(perm))

	if w._accesslogFile != nil {
		w._accesslogFile.Close()
	}
	w._accesslogFile = accessLogFileFd

	if w._accesslog != nil {
		w._accesslog._log.SetOutput(accessLogFileFd)
	} else {
		w._accesslog = &logger{_log: log.New(accessLogFileFd, "", 0), logLevel: LevelDebug}
	}

	return nil
}

func init() {
	err := _fileLogWriter.Init()
	if err != nil {
		fmt.Printf("_fileLogWriter err: %v\n", err)
	}
}

//
//******************access.log文件
//
func AccessInfof(format string, v ...interface{}) {
	err := _fileLogWriter.Accesslog().RawOutput(fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write access info log to file fail, err: %s", err.Error())
	}
}

func AccessErrorf(format string, v ...interface{}) {
	err := _fileLogWriter.Accesslog().RawOutput(fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write access error log to file fail, err: %s", err.Error())
	}
}

func Fatal(s string) {
	err := _fileLogWriter.Log(LevelFatal).Output(LevelFatal, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Fatal log to file fail, err: %s", err.Error())
	}
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelFatal).Output(LevelFatal, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Fatalf log to file fail, err: %s", err.Error())
	}
	os.Exit(1)
}

func Fatalc(ctx context.Context, format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelFatal).Output(LevelFatal, fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Fatalc log to file fail, err: %s", err.Error())
	}
	os.Exit(1)
}

func Error(s string) {
	err := _fileLogWriter.Log(LevelError).Output(LevelError, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Error log to file fail, err: %s", err.Error())
	}
}

func Errorf(format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelError).Output(LevelError, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Errorf log to file fail, err: %s", err.Error())
	}
}

func Errorc(ctx context.Context, format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelError).Output(LevelError, fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Errorc log to file fail, err: %s", err.Error())
	}
}

func Warn(s string) {
	err := _fileLogWriter.Log(LevelWarning).Output(LevelWarning, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Warn log to file fail, err: %s", err.Error())
	}
}

func Warnf(format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelWarning).Output(LevelWarning, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Warnf log to file fail, err: %s", err.Error())
	}
}

func Warnc(ctx context.Context, format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelWarning).Output(LevelWarning, fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Warnc log to file fail, err: %s", err.Error())
	}
}

func Info(s string) {
	err := _fileLogWriter.Log(LevelInfo).Output(LevelInfo, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Info log to file fail, err: %s", err.Error())
	}
}

func Infof(format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelInfo).Output(LevelInfo, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Infof log to file fail, err: %s", err.Error())
	}
}

func Infoc(ctx context.Context, format string, v ...interface{}) {
	str := fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...)
	if len(str) > 300 {
		str = str[:300] + "..."
	}
	err := _fileLogWriter.Log(LevelInfo).Output(LevelInfo, str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Infoc log to file fail, err: %s", err.Error())
	}
}

func Infoc_(ctx context.Context, format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelInfo).Output(LevelInfo, fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Infoc log to file fail, err: %s", err.Error())
	}
}

func Debug(s string) {
	err := _fileLogWriter.Log(LevelDebug).Output(LevelDebug, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Debug log to file fail, err: %s", err.Error())
	}
}

func Debugf(format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelDebug).Output(LevelDebug, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Debugf log to file fail, err: %s", err.Error())
	}
}

func Debugc(ctx context.Context, format string, v ...interface{}) {
	err := _fileLogWriter.Log(LevelDebug).Output(LevelDebug, fmt.Sprintf(fmt.Sprintf("[trace_id=%v] %s", GetTraceId(ctx), format), v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Debugc log to file fail, err: %s", err.Error())
	}
}

func SetLogLevel(level Level) {
	_fileLogWriter.Log(level).SetLogLevel(level)
}

type logger struct {
	_log *log.Logger
	//小于等于该级别的level才会被记录
	logLevel Level
}

// 输出无格式的
func (l *logger) RawOutput(s string) error {
	return l._log.Output(0, s)
}

func (l *logger) Output(level Level, s string) error {
	if l.logLevel < level {
		return nil
	}
	formatStr := "[UNKNOWN] " + s
	switch level {
	case LevelFatal:
		formatStr = "\033[35m[FATAL]\033[0m " + s
	case LevelError:
		formatStr = "\033[31m[ERROR]\033[0m " + s
	case LevelWarning:
		formatStr = "\033[33m[WARN]\033[0m " + s
	case LevelInfo:
		formatStr = "\033[32m[INFO]\033[0m " + s
	case LevelDebug:
		formatStr = "\033[36m[DEBUG]\033[0m " + s
	}
	return l._log.Output(3, formatStr)
}

func (l *logger) Fatal(s string) {
	err := l.Output(LevelFatal, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Fatal log to file fail, err: %s", err.Error())
	}
	os.Exit(1)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	err := l.Output(LevelFatal, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Fatalf log to file fail, err: %s", err.Error())
	}
	os.Exit(1)
}

func (l *logger) Error(s string) {
	err := l.Output(LevelError, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Error log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Errorf(format string, v ...interface{}) {
	err := l.Output(LevelError, fmt.Sprintf(format, v...))
	if err != nil {
		panic("Errorf")
		fmt.Fprintf(os.Stderr, "write Errorf log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Warn(s string) {
	err := l.Output(LevelWarning, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Warn log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Warnf(format string, v ...interface{}) {
	err := l.Output(LevelWarning, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Warnf log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Info(s string) {
	err := l.Output(LevelInfo, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Info log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Infof(format string, v ...interface{}) {
	err := l.Output(LevelInfo, fmt.Sprintf(format, v...))
	if err != nil {
		panic("Errorf")
		fmt.Fprintf(os.Stderr, "write Infof log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Debug(s string) {
	err := l.Output(LevelDebug, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Debug log to file fail, err: %s", err.Error())
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	err := l.Output(LevelDebug, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write Debugf log to file fail, err: %s", err.Error())
	}
}

func (l *logger) SetLogLevel(level Level) {
	l.logLevel = level
}

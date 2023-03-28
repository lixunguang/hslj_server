package request

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"lxg_jz/pkg/logger"
	"net/http"
	"reflect"
	"time"
)

type NewHttpRequestFunc func() (*http.Request, error)
type RequestParams map[string]interface{}

// 设置请求header
type HandlerRequestHeader func(ctx context.Context, header http.Header)

// 设置请求体
type HandlerRequestParams func(ctx context.Context, oldReq interface{}) (newReq interface{})

// 获取server_name 对应的ip以及schema
type ServerHostUrl func(ctx context.Context, serverName string) (string, error)

type Option func(*Options)
type Options struct {
	TimeOut              time.Duration        // 超时时间,默认3s
	RetryMaxNum          int                  // 重试次数，默认重试2次,如果值设置为0则不重试
	RetryIdleTime        time.Duration        // 再次重试的等待时间，默认1ms
	HandlerRequestHeader HandlerRequestHeader // 处理请求头
	HandlerRequestParams HandlerRequestParams // 处理请求参数
	ServerHostUrl        ServerHostUrl        // 请求
	Method               string               //请求方式
}

const (
	defaultMethod        = http.MethodPost //默认Post请求
	defaultTimeOut1s     = 1 * time.Second //默认超时时间1s
	defaultTimeOut2s     = 2 * time.Second
	defaultTimeOut3s     = 3 * time.Second
	defaultTimeOut4s     = 4 * time.Second
	defaultTimeOut5s     = 5 * time.Second
	defaultRetryIdleTime = 1 * time.Millisecond // 默认空闲时间是1ms
)

/**
header or const
*/
const (
	TraceIdHeader  = "X-REQUEST-ID"
	UserIdHeader   = "X-CURRENT-USER-ID"
	UserNameHeader = "X-CURRENT-USER-NAME"
	ServerName     = "server-name"
)

var (
	DefaultTimeOut1sOp = func(op *Options) {
		op.TimeOut = defaultTimeOut1s
	}
	DefaultTimeOut2sOp = func(op *Options) {
		op.TimeOut = defaultTimeOut2s
	}
	DefaultTimeOut3sOp = func(op *Options) {
		op.TimeOut = defaultTimeOut3s
	}
	// 关闭重试
	CloseRetry = func(op *Options) {
		op.RetryMaxNum = -1
	}
	// 默认处理
	DefaultHandlerRequestParams HandlerRequestParams = func(ctx context.Context, oldReq interface{}) (newReq interface{}) {
		//return RequestParams{"params": oldReq}
		return oldReq
	}
	// 默认处理
	DefaultHandlerRequestHandler HandlerRequestHeader = func(ctx context.Context, header http.Header) {
		header.Add("Content-Type", "application/json")
		header.Add(TraceIdHeader, fmt.Sprintf("%v", logger.GetTraceId(ctx)))
	}

	//DefaultServerHost ServerHostUrl = func(ctx context.Context, serverName string) (s string, e error) {
	//	return GetServerUrl(serverName)
	//}
	// 将所有的错误信息，全部放到这里，比如 "error.message" 收集的是{"error":{"message":""}}的信息
	validatorErrorArr = []string{"error.message"}
)

var (
	// error
	validateNilError  = errors.New("the result can not be null")
	validateTypeError = errors.New("the result must be a pointer type")
)

type HttpRequestInfo struct {
	ServiceName string
	Host        string
	Path        string
	Method      string
	Params      interface{}
}

// 这里有个问题是，json 的 unmarshal操作，属于一种懒加载模式，也就是json中没有匹配也不会err，只是没有赋值罢了
func HttpRequestAndDecode(ctx context.Context, serverName string, path string, params interface{}, result interface{}, options ...Option) error {
	if err := validateReceiver(result); err != nil { // result 只支持ptr
		return err
	}
	ctx = context.WithValue(ctx, ServerName, serverName) // clone，不影响原来的ctx，并发安全
	response, err := DefaultHttpRequest(ctx, serverName, path, params, options...)
	if err != nil {
		logger.Errorc(ctx, "[HttpRequest] find err,err=%v", err)
		return err
	}
	isError, errorMsg := CheckJsonHasPattern(response, validatorErrorArr)
	if isError {
		// 这里不记录 error 日志，根据业务需求记录
		return errors.New(errorMsg)
	}
	err = json.Unmarshal(response, result)
	if err != nil {
		logger.Errorc(ctx, "[HttpRequest] json decode http response err: %s", err)
		return fmtError("[HttpRequest] json decode http response err: %s", err)
	}
	return nil
}

// no cached cookie
func HttpRequest(ctx context.Context, request NewHttpRequestFunc, option *Options) ([]byte, error) {
	client := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   option.TimeOut,
	}
	var retryNum = 0
Retry:
	response, err := httpDo(&client, request)
	if err != nil {
		if retryNum < option.RetryMaxNum {
			retryNum++
			logger.Warnc(ctx, "[HttpRequest] request %s err,cur_retry=%d,max_retry=%d,err=%v", getServerName(ctx), retryNum, option.RetryMaxNum, err)
			time.Sleep(option.RetryIdleTime)
			goto Retry
		} else {
			if response != nil && response.Body != nil {
				response.Body.Close()
			}
			return nil, err
		}
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func httpDo(client *http.Client, requestFunc NewHttpRequestFunc) (*http.Response, error) {
	request, err := requestFunc()
	if err != nil {
		return nil, err
	}
	return client.Do(request)
}

func loadOp(option ...Option) *Options {
	op := new(Options)
	for _, elem := range option {
		elem(op)
	}
	if op.HandlerRequestParams == nil {
		op.HandlerRequestParams = DefaultHandlerRequestParams
	}
	if op.HandlerRequestHeader == nil {
		op.HandlerRequestHeader = DefaultHandlerRequestHandler
	}
	if op.TimeOut == 0 {
		op.TimeOut = defaultTimeOut3s
	}
	if op.RetryIdleTime == 0 {
		op.RetryIdleTime = defaultRetryIdleTime
	}

	if op.Method == "" {
		op.Method = defaultMethod
	}
	return op
}

func handlerRequestHeader(ctx context.Context, req *http.Request, op *Options) {
	if op != nil && op.HandlerRequestHeader != nil {
		op.HandlerRequestHeader(ctx, req.Header)
	}
}

func handlerRequestParams(ctx context.Context, req interface{}, op *Options) (io.Reader, error) {
	var buf bytes.Buffer
	if op != nil && op.HandlerRequestParams != nil {
		req = op.HandlerRequestParams(ctx, req)
	}
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return nil, fmtError("[AddJsonRequestParams] err,err=%+v,req=%+v", err, req)
	}
	return ioutil.NopCloser(&buf), nil
}

func fmtError(format string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(format, v...))
}

func validateReceiver(result interface{}) error {
	if result == nil {
		return validateNilError
	} else {
		kind := reflect.TypeOf(result).Kind()
		if kind != reflect.Ptr {
			return validateTypeError
		}
		return nil
	}
}

func getServerName(ctx context.Context) string {
	s, _ := ctx.Value(ServerName).(string)
	return s
}

func DefaultHttpRequest(ctx context.Context, host string, path string, params interface{}, options ...Option) ([]byte, error) {
	beginTime := time.Now()
	option := loadOp(options...)
	result, err := HttpRequest(ctx, func() (*http.Request, error) {
		request, err := NewProxyRequest(&HttpRequestInfo{
			Host:   host,
			Method: option.Method,
			Params: params,
			Path:   path,
		}, ctx, option)
		if err != nil {
			return nil, err
		}
		logger.Infoc(ctx, "[HttpRequest] start,url=%s,host=%s,params=%+v", host, request.URL, params)
		return request, nil
	}, option)
	if err != nil {
		return nil, err
	}
	logger.Infoc(ctx, "[HttpRequest] end,host=%s,path=%s,spend=%fs,result=%s", host, path, time.Now().Sub(beginTime).Seconds(), result)
	return result, err
}

// 内部代理的http的request
func NewProxyRequest(requestInfo *HttpRequestInfo, ctx context.Context, op *Options) (*http.Request, error) {
	if requestInfo.Host == "" {
		return nil, fmtError("[NewProxyRequest] url err, url is empty")
	}
	// json请求参数
	reader, err := handlerRequestParams(ctx, requestInfo.Params, op)
	if err != nil {
		return nil, err
	}
	// 创建request
	request, err := http.NewRequest(requestInfo.Method, requestInfo.Host+requestInfo.Path, reader)
	if err != nil {
		return nil, fmtError("[NewRequest] err,err=%+v,host=%s,path=%s,params=%+v", err, requestInfo.Host, requestInfo.Path, requestInfo.Params)
	}
	// 请求头
	handlerRequestHeader(ctx, request, op)

	return request, nil
}

// Package common provides commonly used functions and utilities.
package common

import (
	"hslj/config"
	"fmt"
)

var SigKey = []byte("signature_hmac_secret_shared_key")

const (
	Failed  = -1
	Success = 0
)

const (
	FailedID = -1 //-1为不合法的id，一般为插入失败,最初为0  //todo：重构
)

//错误处理
var ()

const (
	NotExpired       = 0 //是否过期
	TokenExpired     = 1
	TokenFormatError = 2 //
)

const (
	PictureType = 1  //图片资源类型
	WordType    = 2  //word文档资源类型
	PdfType     = 3  //pdf资源类型
	AppType     = 4  //app资源类型
	IbeType     = 5  //ibe资源类型
	MessageType = 6  //同7,未使用
	RichType    = 7  //富文本资源类型
	VedioType   = 8  //视频资源类型
	FileType    = 23 //除了MessageType和RichType之外的所有类型
)

//课程资源子类型
const (
	CoursePicture = 1 //课程缩略图
	CourseVedio   = 2 //课程介绍视频
	CourseLesson  = 3 //课程章节
)

//章节子类型
const (
	LessonContent    = 1
	LessonRefer      = 2
	LessonExperiment = 3
	LessonWork       = 4
)

//作业资源类型
const (
	WorkRequireRichText   = 41 //作业要求：富文本格式要求
	WorkRequireAttachment = 42 //作业要求：文件附件
	WorkCommitAttachment  = 43 //作业提交的文件
	WorkCommitText        = 44 //作业提交附言
)

const (
	StudentRole = 1 //人员类型；学生1、教师2、管理员3
	TeacherRole = 2
	AdminRole   = 3
)

// GetCDNAddr is a function that returns the CDN address.
func GetCDNAddr() string {
	addr := fmt.Sprintf("%s", config.Vipper.Get("CDNServer.Addr"))

	return addr
}

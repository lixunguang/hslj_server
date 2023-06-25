package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

//获取课程列表
func CourseAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseAll Controller")
	//获取参数

	//参数校验

	//调用service
	var courseAllRes []admin_dto.CourseAllRes
	var err cerror.Cerror
	courseAllRes, err = admin_service.CourseALL(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, courseAllRes)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseAll Controller", courseAllRes, err)
}

//删除课程
func CourseDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseDel Controller")
	//获取参数
	var param admin_dto.CourseIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseDel Controller", param)
	}
	//参数校验

	//调用service
	id, err := admin_service.CourseDel(ctx, param.ID)

	//结果返回
	if id == common.FailedID {
		util.FailJson(ctx, err)
	} else {
		res := dto.IDRes{ID: id}
		util.SuccessJson(ctx, res)
	}

	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "CourseDel Controller", id, err)
}

// 获取课程的信息
func Course(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "Course Controller")
	//获取参数
	var param admin_dto.CourseIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "Course Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.Course(ctx, param.ID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "Course Controller", res, err)
}

//增加一门课程
func CourseAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseAdd Controller")
	//获取参数
	var param admin_dto.AddCourseAllParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseAdd Controller", param)
	}
	//参数校验

	//调用service
	var param1 admin_dto.AddCourseParam
	param1.Title = param.Title
	param1.Desc = param.Desc

	id, err1 := admin_service.AddCourse(ctx, param1)
	//结果返回
	/*
	   if  err1!= nil {
	   		util.FailJson(ctx, err1)
	   	} else {
	   		res := dto.IDRes{ID: id}
	   		util.SuccessJson(ctx, res)
	   	}
	*/
	if id >= 1 && err1 == nil {
		var param2 admin_dto.CourseExtraParam
		param2.CourseID = id
		param2.VedioID = param.VedioID
		param2.PictureID = param.PictureID
		param2.AuthorID = param.AuthorID

		_, err2 := admin_service.AddCourseExtra(ctx, param2)
		//结果返回
		if err2 != nil {
			util.FailJson(ctx, err2)
		} else {
			res := dto.IDRes{ID: id}
			util.SuccessJson(ctx, res)
		}
	} else {
		util.FailJson(ctx, cerror.ErrorAddCourseFailed)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseAdd Controller", id, err1)
}

//增加课程的章节：即增加了章节，又把章节与课程相关联
func CourseLessonAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseLessonAdd Controller")
	//获取参数
	var param admin_dto.AddCourseLessonParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseLessonAdd Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.AddLesson(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseLessonAdd Controller", res, err)
}

//获取课程的所有章节
func CourseLessonAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseLessonAll Controller")
	//获取参数
	var param admin_dto.CourseIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseLessonAll Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.CourseLessonAll(ctx, param.ID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseLessonAll Controller", res, err)
}

func CourseLessonDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseLessonDel Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseLessonDel Controller", param)
	}
	//参数校验

	//调用service
	err := admin_service.CourseLessonDel(ctx, param.CourseID, param.LessonID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		res := dto.IDRes{ID: param.LessonID}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "CourseLessonDel Controller", param.LessonID, err)
}

// 更新基本信息
func CourseUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseUpdate Controller")
	//获取参数
	var param admin_dto.UpdateCourse
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseUpdate Controller", param)
	}
	//参数校验

	//调用service
	id, err := admin_service.CourseUpdate(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var res = dto.IDParam{ID: id}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "CourseUpdate Controller", id, err)
}

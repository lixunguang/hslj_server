package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddUser Controller")
	//获取参数
	var param admin_dto.User
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddUser Controller", param)
	}

	if param.LoginID == "" {
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	//参数校验

	//调用service
	res, err := admin_service.AddUser(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AddUser Controller", res, err)
}

func DelUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "DelUser Controller")
	//获取参数
	var param admin_dto.IDParam2
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "DelUser Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelUser(ctx, param.LoginID)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "DelUser Controller", res, err)
}

func UpdateUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "UpdateUser Controller")
	//获取参数
	var param admin_dto.User
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "UpdateUser Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.UpdateUser(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "UpdateUser Controller", res, err)
}

func AllUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AllUser Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := admin_service.AllUser(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AllUser Controller", res, err)
}

func AddAdmin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddAdmin Controller")
	//获取参数
	var param admin_dto.Admin
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddAdmin Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.AddAdmin(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AddAdmin Controller", res, err)
}

func DelAdmin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "DelAdmin Controller")

	//获取参数
	var param admin_dto.AdminParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "DelAdmin Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelAdmin(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "DelAdmin Controller", res, err)
}

//暂时写在这个文件
func AddTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddTeacher Controller")
	//获取参数
	var param admin_dto.AddTeacherParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddTeacher Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.AddTeacher(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AddTeacher Controller", res, err)
}

//暂时写在这个文件
func AllTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AllTeacher Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := admin_service.AllTeacher(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "AllTeacher Controller", res, err)
}

//暂时写在这个文件
func DelTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "DelTeacher Controller")
	//获取参数
	var param admin_dto.IDParam2
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "DelTeacher Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelTeacher(ctx, param.LoginID)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "DelTeacher Controller", res, err)
}

//暂时写在这个文件
func UpdateTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "UpdateTeacher Controller")
	//获取参数
	var param admin_dto.UpdateTeacherParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "UpdateTeacher Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.UpdateTeacher(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "UpdateTeacher Controller", res, err)
}

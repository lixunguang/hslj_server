package controller

import (
<<<<<<< HEAD
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func AddUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddUser Controller")
	//获取参数
	var param dto.User
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
	res, err := service.AddUser(ctx, param)

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
	var param dto.IDParam2
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
	res, err := service.DelUser(ctx, param.LoginID)

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
	var param dto.User
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
	res, err := service.UpdateUser(ctx, param)

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
	res, err := service.AllUser(ctx)

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
	var param dto.Admin
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
	res, err := service.AddAdmin(ctx, param)

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
	var param dto.AdminParam
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
	res, err := service.DelAdmin(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "DelAdmin Controller", res, err)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func GetActionAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetActionAll Controller")

	// 获取参数

	//参数校验

	//调用service

	res, err := service.ActionAll(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetActionAll Controller", res, "tokenStr")
}

func GetAction(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetAction Controller")

	// 获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetAction Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.ActionDetail(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetAction Controller", res, "tokenStr")
}

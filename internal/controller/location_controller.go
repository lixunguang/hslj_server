package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func GetLocationList(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetLocationList Controller")

	// 获取参数
	var param dto.GetLocationListParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetLocationList Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.GetLocation(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetLocationList Controller", res, "tokenStr")
}

func AddLocation(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddLocation Controller")
	// 获取参数
	var param dto.AddLocationParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddLocation Controller", param)
	}
	//参数校验

	//调用service
	var res dto.AddLocationRes

	var err cerror.Cerror = nil

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "AddLocation Controller", res, "tokenStr")
}

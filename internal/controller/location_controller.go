package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func GetLocation(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetAreaLocationList Controller")

	// 获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetAreaLocationList Controller", param)
	}

	//参数校验

	//调用service
	res, err := service.GetLocation(ctx, param.ID)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetAreaLocationList Controller", res, "tokenStr")
}

func GetAreaLocationList(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetAreaLocationList Controller")

	// 获取参数
	var param dto.GetLocationListParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetAreaLocationList Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.GetAreaLocationList(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetAreaLocationList Controller", res, "tokenStr")
}

func GetLocationActive(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetLocationActive Controller")

	// 获取参数
	var param dto.GetLocationListParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetLocationActive Controller", param)
	}

	//参数校验

	//调用service
	res, err := service.GetAreaLocationList(ctx, param) //todo:需要修改

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetLocationActive Controller", res, "tokenStr")
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
	param.Rating = 4
	param.ISAuth = 0

	res, err := service.AddLocation(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "AddLocation Controller", res, "tokenStr")
}

func GetLocationHot(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetLocationHot Controller")

	// 获取参数

	//参数校验

	//调用service

	res, err := service.GetLocationHot(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetLocationHot Controller", res, "tokenStr")
}

package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func GetNewsLatest(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetNewsLatest Controller")

	// 获取参数

	//参数校验

	//调用service
	res, err := service.NewsLatest(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetNewsLatest Controller", res, "tokenStr")
}

func GetNewsBanner(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetNewsLatest Controller")

	// 获取参数

	//参数校验

	//调用service
	res, err := service.NewsBanner(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetNewsLatest Controller", res, "tokenStr")
}

func GetNewsAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetNewsAll Controller")

	// 获取参数

	//参数校验

	//调用service

	res, err := service.NewsALL(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetNewsAll Controller", res, "tokenStr")
}

func GetNews(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetNews Controller")

	// 获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "GetNews Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.NewsDetail(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetNews Controller", res, "tokenStr")
}

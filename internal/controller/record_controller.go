package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

func AddRecord(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddRecord Controller")

	// 获取参数
	var param dto.AddRecordParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddRecord Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.AddRecord(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, dto.IDRes{ID:res})
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "AddRecord Controller", res, "tokenStr")
}


func GetRecordByUserID(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddRecord Controller")

	// 获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddRecord Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.GetRecordByUserID(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "AddRecord Controller", res, "tokenStr")
}

func GetRecordByLocationID(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "AddRecord Controller")

	// 获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddRecord Controller", param)
	}

	//参数校验

	//调用service

	res, err := service.GetRecordByLocationID(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "AddRecord Controller", res, "tokenStr")
}


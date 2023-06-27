package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
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
		logger.Infoc(ctx, "[%s] input param: %+v", "Login Controller", param)
	}

	//参数校验

	//调用service

	var res []dto.GetLocationListRes

	var item dto.GetLocationListRes
	item.Name = "北京颐和园"
	item.Desc = "颐和园描述。。。"
	item.Latitude = 1.11
	item.Longitude = 2.22
	item.Rating = 4
	//item.OpenTime =

	res = append(res, item)

	item.Name = "北京天坛"
	item.Desc = "天坛描述。。。"
	item.Latitude = 3.3333
	item.Longitude = 4.4444
	item.Rating = 4
	res = append(res, item)

	var err cerror.Cerror = nil
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
		logger.Infoc(ctx, "[%s] input param: %+v", "Login Controller", param)
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

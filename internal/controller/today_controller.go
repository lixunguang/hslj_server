package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/service"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

//today详细信息
func GetToday(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetToday Controller")

	// 获取参数

	//参数校验

	//调用service
	res, err := service.GetToday(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetToday Controller", res, "tokenStr")
}

//主页显示的today信息
func GetTodayShort(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetTodayShort Controller")

	// 获取参数

	//参数校验

	//调用service
	res, err := service.GetTodayShort(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "GetTodayShort Controller", res, "tokenStr")
}

// Package controller is ...
package admin_controller

import (
	"edu-imp/internal/admin_service"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

func News(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "News Controller")
	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "News Controller", param)
	}
	//参数校验

	//调用service
	news, err := admin_service.OneNews(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)

	} else {
		util.SuccessJson(ctx, news)
	}
	logger.Infoc(ctx, "[%s] end***  result:  news=%+v,err=%+v", "News Controller", news, err)
}

// NewsAdd is 增加新闻.
func NewsAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "NewsAdd Controller")
	// 获取参数
	var param dto.NewsAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "NewsAdd Controller", param)
	}
	// 参数校验

	// 调用service
	res, err := admin_service.AddNews(ctx, param)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "NewsAdd Controller", res, err)
}

// NewsDelete is 刪除新闻.
func NewsDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "NewsDelete Controller")
	// 获取参数
	var param dto.DelNewsParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "NewsDel Controller", param)
	}
	// 参数校验

	// 调用service
	res, err := admin_service.DelNews(ctx, param.ID)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "NewsDel Controller", res, err)
}

// 编辑新闻
func NewsUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "NewsUpdate Controller")
	// 获取参数
	var param dto.NewsUpdateParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "NewsUpdate Controller", param)
	}
	// 参数校验

	// 调用service
	res, err := admin_service.UpdateNews(ctx, param)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "NewsUpdate Controller", res, err)
}

func NewsTitleAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "NewsTitleALL Controller")
	//获取参数

	//参数校验

	//调用service

	res, err := admin_service.NewsTitleALL(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "NewsTitleAll Controller", res, err)
}

func NewsBannerAll(ctx *gin.Context) {

	// 获取参数

	// 参数校验

	// 调用service
	res, err := service.BannerNewsALL(ctx)

	// 结果返回

	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "NewsBannerAll Controller", res, err)
}

package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddRichtext(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "ResourceAdd Controller")
	//获取参数
	var param admin_dto.ResourceParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "AddRichtext Controller", param)
	}

	var param2 dto.Resource
	param2.Title = param.Title
	param2.Desc = param.Desc
	param2.Content = param.Content
	param2.Type = common.RichType
	// 调用service
	id, err := service.AddResource(ctx, param2)
	// 结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var res = dto.IDParam{ID: id}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "AddRichtext Controller", id, err)
}

func DelRichtext(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "DelRichtext Controller")

	//获取参数
	var param admin_dto.ResouceIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "DelRichtext Controller", param)
	}

	// 调用service
	var param2 admin_dto.ResouceDelParam
	param2.ID = param.ID
	param2.Type = common.RichType
	id, err := service.DelResource(ctx, param2)

	// 结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var res = dto.IDParam{ID: id}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "DelRichtext Controller", id, err)
}

func UpdateRichtext(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "UpdateRichtext Controller")
	//获取参数
	var param dto.UpdateResourceParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "UpdateRichtext Controller", param)
	}

	param.Type = common.RichType
	// 调用service
	id, err := service.UpdateResource(ctx, param)
	// 结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var res = dto.IDParam{ID: id}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "UpdateRichtext Controller", id, err)
}

func DelFile(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "DelFile Controller")
	//获取参数
	var param admin_dto.ResouceIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "DelFile Controller", param)
	}

	// 调用service
	var param2 admin_dto.ResouceDelParam
	param2.ID = param.ID
	param2.Type = common.FileType
	id, err := service.DelResource(ctx, param2)

	// 结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var res = dto.IDParam{ID: id}
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  id=%+v,err=%+v", "DelFile Controller", id, err)
}

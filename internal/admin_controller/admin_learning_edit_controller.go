package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

// LearningAll 获取所有学习资源
func LearningAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningAll Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := admin_service.GetLearningAll(ctx)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningAll Controller", res, err)
}

//获取学习资源详情（根据id）
func Learning(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "Learning Controller")
	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v,param:%+v", "method", err, ctx.Request)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "Learning Controller", param)
	}
	//参数校验
	//调用service
	res := admin_service.GetLearningByID(ctx, param.ID)

	//结果返回

	util.SuccessJson(ctx, res)
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v", "Learning Controller", res)
}

// LearningAdd 增加学习资源
func LearningAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningAdd Controller")
	//获取参数
	var param dto.Learning
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningAdd Controller", param)
	}
	//参数校验
	//param.Author = "云道教育" //todo：和新闻统一考虑
	//调用service
	res, err := admin_service.AddLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningAdd Controller", res, err)
}

// LearningDel 删除学习资源
func LearningDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningDel Controller")
	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningDel Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningDel Controller", res, err)
}

// LearningUpdate 更新学习资源
func LearningUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningUpdate Controller")
	//获取参数
	var param admin_dto.UpdateLearning
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningUpdate Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.UpdateLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDParam{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningUpdate Controller", res, err)
}

//获取所有分类
func LearningCategory(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningCategory Controller")
	//获取参数

	//参数校验
	//调用service
	res, err := admin_service.GetLearningCategory(ctx)

	//结果返回

	util.SuccessJson(ctx, res)
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningCategory Controller", res, err)
}

//增加学习资源分类
func LearningCategoryAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningCategoryAdd Controller")

	//获取参数
	var param dto.LearningCategory
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningCategoryAdd Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.AddLearningCategory(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningCategoryAdd Controller", res, err)
}

//增加学习资源分类
func LearningCategoryUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningCategoryUpdate Controller")

	//获取参数
	var param dto.UpdateLearningCategoryParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningCategoryUpdate Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.UpdateLearningCategory(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningCategoryUpdate Controller", res, err)
}

//删除学习资源分类
func LearningCategoryDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningCategoryDel Controller")

	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningCategoryDel Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelLearningCategory(ctx, param.ID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningCategoryDel Controller", res, err)
}

// LearningResource 获取学习资源的视频资源.
func LearningResource(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningResource Controller")

	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningResource Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.GetLearningResource(ctx, param.ID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningResource Controller", res, err)
}

// LearningResourceAdd 增加资源条目.
func LearningResourceAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningResourceAdd Controller")

	//获取参数
	var param admin_dto.AddLearningResourceParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningResourceAdd Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.AddLearningResource(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningResourceAdd Controller", res, err)
}

// LearningResourceDel 增加资源条目.
func LearningResourceDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LearningResourceDel Controller")

	//获取参数
	var param admin_dto.DelLearningResourceParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LearningResourceDel Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.DelLearningResource(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LearningResourceDel Controller", res, err)
}

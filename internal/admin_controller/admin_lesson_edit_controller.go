package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/internal/common"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

//获取课程详情
func CourseLesson(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseLesson Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseLesson Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.CourseLesson(ctx, param.CourseID, param.LessonID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {

		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseLesson Controller", res, err)
}

func CourseLessonUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "CourseLessonUpdate Controller")
	//获取参数
	var param admin_dto.UpdateLessonParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "CourseLessonUpdate Controller", param)
	}
	//参数校验

	//调用service
	res, err := admin_service.LessonUpdate(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "CourseLessonUpdate Controller", res, err)
}

//根据lessonid 获取课程内容
func LessonContent(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonContent Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonContent Controller", param)
	}
	//参数校验

	//调用service
	var sectionType int
	sectionType = common.LessonContent //内容类型
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonContent Controller", res, err)
}

func LessonContentAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonContentAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonContentAdd Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonContentAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonContentAdd Controller", res, err)
}

func LessonContentDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonContentDel Controller")
	//获取参数
	var param admin_dto.LessonContentParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonContentDel Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonContentDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonContentDel Controller", res, err)
}

//根据lessonid 获取课程参考资源
func LessonRefer(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonRefer Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonRefer Controller", param)
	}
	//参数校验

	//调用service
	var sectionType int
	sectionType = common.LessonRefer //内容类型
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonRefer Controller", res, err)
}

func LessonReferAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonReferAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonReferAdd Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonReferAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonReferAdd Controller", res, err)
}

func LessonReferDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonReferDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonReferDel Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonReferDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonReferDel Controller", res, err)
}

func LessonExperiment(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonExperiment Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonExperiment Controller", param)
	}
	//参数校验

	//调用service
	var sectionType int
	sectionType = common.LessonExperiment //内容类型
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonExperiment Controller", res, err)
}

func LessonExperimentAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonExperimentAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonExperimentAdd Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonExperimentAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonExperimentAdd Controller", res, err)
}

func LessonExperimentDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonExperimentDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonExperimentDel Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonExperimentDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonExperimentDel Controller", res, err)
}

func LessonWork(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonExperiment Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonWork Controller", param)
	}
	//参数校验

	//调用service
	var sectionType int
	sectionType = common.LessonWork //内容类型
	ret, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	var res []admin_dto.LessonSectionContent
	for _, val := range ret {
		if val.Index != common.WorkRequireRichText {
			res = append(res, val)
		}
	}

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonWork Controller", res, err)
}

func LessonWorkAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonWorkAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonWorkAdd Controller", param)
	}
	param.ResourceIndex = common.WorkRequireAttachment
	//参数校验
	res, err := admin_service.LessonWorkAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonWorkAdd Controller", res, err)
}

func LessonWorkDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonExperimentDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonWorkDel Controller", param)
	}

	//参数校验
	res, err := admin_service.LessonWorkDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonWorkDel Controller", res, err)
}

func LessonWorkRequirement(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonWorkRequirement Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonWorkRequirement Controller", param)
	}
	//参数校验

	//调用service
	var sectionType int
	sectionType = common.LessonWork //内容类型

	resource, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	var res admin_dto.LessonSectionContent
	for _, val := range resource {
		if val.Index == common.WorkRequireRichText {
			res = val
		}
	}
	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonWorkRequirement Controller", res, err)
}

func LessonWorkRequirementUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "LessonWorkRequirementUpdate Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "LessonWorkRequirementUpdate Controller", param)
	}

	//参数校验
	param.ResourceIndex = common.WorkRequireRichText
	res, err := admin_service.LessonWorkRequirementUpdate(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "LessonWorkRequirementUpdate Controller", res, err)
}

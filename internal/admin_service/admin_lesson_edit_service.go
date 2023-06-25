package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 增加课程-章节
func AddLesson(ctx *gin.Context, param admin_dto.AddCourseLessonParam) (admin_dto.AddLessonRes, cerror.Cerror) {
	var res admin_dto.AddLessonRes

	var lessonParam dto.Lesson

	//add lesson
	lessonParam.Title = param.LessonTitle
	lessonParam.Desc = param.LessonDesc
	lesson, _ := dao.AddLesson(ctx, lessonParam)

	//course_resource
	var lessonResourceParam admin_dto.CourseResourceParam
	lessonResourceParam.CourseID = param.CourseID
	lessonResourceParam.ResourceID = lesson.ID
	lessonResourceParam.Type = common.CourseLesson
	lessonResourceParam.Index = 0

	dao.AddCourseResource(ctx, lessonResourceParam)

	//默认每个章节有包含4部分内容
	var lessonSectionParam dto.LessonSection
	lessonSectionParam.LessonID = lesson.ID

	lessonSectionParam.Title = "教学内容"
	lessonSectionParam.Index = 1
	dao.AddLessonSection(ctx, lessonSectionParam)

	lessonSectionParam.Title = "参考资源"
	lessonSectionParam.Index = 2
	dao.AddLessonSection(ctx, lessonSectionParam)

	lessonSectionParam.Title = "做实验"
	lessonSectionParam.Index = 3
	dao.AddLessonSection(ctx, lessonSectionParam)

	lessonSectionParam.Title = "作业"
	lessonSectionParam.Index = 4
	dao.AddLessonSection(ctx, lessonSectionParam)

	res.CourseId = param.CourseID
	res.LessonId = lesson.ID
	return res, nil
}

//获取Lesson详情
func CourseLesson(ctx *gin.Context, CourseID int, lessonID int) (admin_dto.LessonItem, cerror.Cerror) {

	var res admin_dto.LessonItem

	lesson, err := dao.GetLesson(ctx, lessonID)
	if err != nil {
		return res, err
	}

	res.ID = lesson.ID
	res.Title = lesson.Title
	res.Desc = lesson.Desc

	return res, nil

}

func LessonUpdate(ctx *gin.Context, param admin_dto.UpdateLessonParam) (int, cerror.Cerror) {

	return dao.LessonUpdate(ctx, param.CourseID, param.LessonID, param.Title, param.Desc)

}

func GetCourseLessonSectionResource(ctx *gin.Context, lessonID int, sectionType int) ([]admin_dto.LessonSectionContent, cerror.Cerror) {

	//根据lessonid 和type 获取lesson_section 的id,返回值为一个id
	lessonSectionID, err := dao.GetLessonSectionIDByLessonIDType(ctx, lessonID, sectionType)
	if err != nil {
		return nil, err
	}

	if lessonSectionID == 0 { //可以写的更严谨一点，小于等于0均为无效id
		return nil, err
	}

	var LessonSectionContentList []admin_dto.LessonSectionContent
	//通过lesson_section_id 获取资源的id （一般为多个）
	lessonSectionResource, err2 := dao.GetLessonSectionResourceBySectionIDType(ctx, lessonSectionID)
	if err2 != nil {
		return nil, err2
	}

	for _, val := range lessonSectionResource {
		var item admin_dto.LessonSectionContent

		item.LessonID = lessonID
		item.ResourceID = val.ResourceID
		item.Title, _ = dao.GetResourceTitleByID(ctx, val.ResourceID)
		item.Desc, _ = dao.GetResourceDescByID(ctx, val.ResourceID)

		str := dao.GetResourceContentFromID(ctx, val.ResourceID)
		if len(str) > 100 {
			str = str[:100]
			str += " ..."
		}

		item.Content = str
		contentType, _ := dao.GetResourceTypeByID(ctx, val.ResourceID)
		item.ContentType = int(contentType)
		item.Index = val.Index

		LessonSectionContentList = append(LessonSectionContentList, item)
	}

	//根据资源的id获取资源详情
	return LessonSectionContentList, nil

}

func LessonContentAdd(ctx *gin.Context, param admin_dto.LessonResourceAddParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID
	resourceParam.Index = param.ResourceIndex

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonContent)
	if lessonSectionID == 0 || err2 != nil {
		return dto.IDRes{ID: 0}, cerror.ErrorAddLessonContentFind
	}

	resourceParam.SectionID = lessonSectionID
	res, err := dao.AddLessonSectionResource(ctx, resourceParam)
	return res, err

}

func LessonContentDel(ctx *gin.Context, param admin_dto.LessonContentParam) (dto.IDParam, cerror.Cerror) {

	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonContent)
	if err2 != nil {
		return dto.IDParam{ID: 0}, err2
	}

	var delParam dto.LessonSectionDelResource
	delParam.ResourceID = param.ResourceID

	delParam.SectionID = lessonSectionID
	deleteLine1, err := dao.DelLessonSectionResource(ctx, delParam)
	fmt.Println(deleteLine1)

	var res = dto.IDParam{ID: param.ResourceID}
	return res, err
}

func LessonReferAdd(ctx *gin.Context, param admin_dto.LessonResourceAddParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID
	resourceParam.Index = param.ResourceIndex

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonRefer)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}
	resourceParam.SectionID = lessonSectionID

	res, err := dao.AddLessonSectionResource(ctx, resourceParam)

	return res, err
}

func LessonReferDel(ctx *gin.Context, param admin_dto.LessonResourceDelParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionDelResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonRefer)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}

	resourceParam.SectionID = lessonSectionID
	lines, err := dao.DelLessonSectionResource(ctx, resourceParam)
	fmt.Println(lines)

	return dto.IDRes{ID: param.ResourceID}, err
}

func LessonExperimentAdd(ctx *gin.Context, param admin_dto.LessonResourceAddParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID
	resourceParam.Index = param.ResourceIndex

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonExperiment)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}
	resourceParam.SectionID = lessonSectionID

	res, err := dao.AddLessonSectionResource(ctx, resourceParam)

	return res, err
}

func LessonExperimentDel(ctx *gin.Context, param admin_dto.LessonResourceDelParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionDelResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonExperiment)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}

	resourceParam.SectionID = lessonSectionID
	lines, err := dao.DelLessonSectionResource(ctx, resourceParam)
	fmt.Println(lines)

	//todo：删除最终的富文本或者文件

	return dto.IDRes{ID: param.ResourceID}, err
}

func LessonWorkRequirementUpdate(ctx *gin.Context, param admin_dto.LessonResourceAddParam) (dto.IDRes, cerror.Cerror) {

	var sectionType = common.LessonWork //内容类型
	allRes, _ := GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)
	var content admin_dto.LessonSectionContent
	for _, val := range allRes {
		if val.Index == common.WorkRequireRichText {
			content = val
		}
	}

	//先删除已有的
	var delParam admin_dto.LessonResourceDelParam
	delParam.CourseID = param.CourseID
	delParam.LessonID = param.LessonID
	delParam.ResourceID = content.ResourceID
	delRes, delErr := LessonWorkDel(ctx, delParam)
	if delErr != nil {
		return delRes, delErr
	}

	param.ResourceIndex = common.WorkRequireRichText
	//增加新的
	addRes, addErr := LessonWorkAdd(ctx, param)
	if addErr != nil {
		return addRes, addErr
	}

	return addRes, nil
}

func LessonWorkAdd(ctx *gin.Context, param admin_dto.LessonResourceAddParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID
	resourceParam.Index = param.ResourceIndex //common.WorkRequireRichText //param.ResourceIndex//？？？？这里实现错了，目前这里是增加requirement

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonWork)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}
	resourceParam.SectionID = lessonSectionID

	res, err := dao.AddLessonSectionResource(ctx, resourceParam)

	return res, err
}

func LessonWorkDel(ctx *gin.Context, param admin_dto.LessonResourceDelParam) (dto.IDRes, cerror.Cerror) {

	var resourceParam dto.LessonSectionDelResource
	//在lesson_section_resource表里增加一项
	resourceParam.ResourceID = param.ResourceID

	//查询lessonsectionid
	lessonSectionID, err2 := dao.GetLessonSectionIDByLessonIDType(ctx, param.LessonID, common.LessonWork)
	if err2 != nil {
		return dto.IDRes{ID: 0}, err2
	}

	resourceParam.SectionID = lessonSectionID
	lines, err := dao.DelLessonSectionResource(ctx, resourceParam)
	fmt.Println(lines)

	return dto.IDRes{ID: param.ResourceID}, err
}

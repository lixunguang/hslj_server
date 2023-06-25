package service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func LessonAdd(ctx *gin.Context, param dto.Lesson) (dto.IDRes, cerror.Cerror) {

	return dao.AddLesson(ctx, param)

}

func LessonSectionResourceAdd(ctx *gin.Context, param dto.LessonSectionResource) (dto.IDRes, cerror.Cerror) {

	return dao.AddLessonSectionResource(ctx, param)

}

func CourseLessonOverview(ctx *gin.Context, param dto.IDParam) (dto.LessonOverviewRes, cerror.Cerror) {
	var res dto.LessonOverviewRes

	var err cerror.Cerror
	res.LessonTitle, err = dao.GetLessonTitleByID(ctx, param.ID)
	if err != nil {
		return res, err
	}

	var lessonSection []dto.LessonSection
	lessonSection, err = dao.GetLessonSectionByLessonID(ctx, param.ID)
	if err != nil {
		return res, err
	}

	res.LessonSections = lessonSection

	return res, nil
}

func GetCourseLessonSectionResource(ctx *gin.Context, lessonID int, sectionType int) ([]dto.LessonSectionContent, cerror.Cerror) {

	//根据lessonid 和type 获取lesson_section 的id,返回值为一个id
	lessonSectionID, err := dao.GetLessonSectionIDByLessonIDType(ctx, lessonID, sectionType)
	if err != nil {
		return nil, err
	}

	if lessonSectionID == 0 { //可以写的更严谨一点，小于等于0均为无效id
		return nil, err
	}

	var LessonSectionContent []dto.LessonSectionContent
	//通过lesson_section_id 获取资源的id （一般为多个）
	lessonSectionResource, err2 := dao.GetLessonSectionResourceBySectionIDType(ctx, lessonSectionID)
	if err2 != nil {
		return nil, err2
	}

	for _, val := range lessonSectionResource {
		var item dto.LessonSectionContent

		item.LessonID = lessonID
		item.Title, _ = dao.GetResourceTitleByID(ctx, val.ResourceID)
		item.Index = val.Index
		item.Content = dao.GetResourceContentFromID(ctx, val.ResourceID)
		contentType, _ := dao.GetResourceTypeByID(ctx, val.ResourceID)
		item.ContentType = int(contentType)

		LessonSectionContent = append(LessonSectionContent, item)
	}

	//根据资源的id获取资源详情
	return LessonSectionContent, nil

}

func CourseLessonSectionWorkResource(ctx *gin.Context, lessonID int, sectionType int) (dto.LessonSectionContentWork, cerror.Cerror) {
	var res dto.LessonSectionContentWork

	//根据lessonid 和type 获取lesson_section 的id,返回值为一个id
	lessonSectionID, err := dao.GetLessonSectionIDByLessonIDType(ctx, lessonID, sectionType)
	if err != nil {
		return res, err
	}

	if lessonSectionID == 0 { //可以写的更严谨一点，小于等于0均为无效id
		return res, err
	}

	//通过lesson_section_id 获取资源的id （一般为多个）
	lessonSectionResource, err2 := dao.GetLessonSectionResourceBySectionIDType(ctx, lessonSectionID)
	if err2 != nil {
		return res, err2
	}

	for _, val := range lessonSectionResource {

		title, _ := dao.GetResourceTitleByID(ctx, val.ResourceID)
		//index := val.Index
		content := dao.GetResourceContentFromID(ctx, val.ResourceID)
		contentType, _ := dao.GetResourceTypeByID(ctx, val.ResourceID)

		//给res赋值，转换下数据类型
		if val.Index == common.WorkRequireRichText {
			res.Requirement = content
		} else if val.Index == common.WorkRequireAttachment {
			var workItem dto.WorkContentItem
			workItem.Title = title
			workItem.Content = content
			workItem.ContentType = int(contentType)

			res.Contents = append(res.Contents, workItem)
		}

	}

	res.LessonID = lessonID

	//根据资源的id获取资源详情
	return res, nil

}

func CourseLessonWork(ctx *gin.Context, id int) {

}

func CourseLessonModify(ctx *gin.Context) {

}

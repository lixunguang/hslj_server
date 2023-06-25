package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CourseALL(ctx *gin.Context) ([]admin_dto.CourseAllRes, cerror.Cerror) {

	var res []admin_dto.CourseAllRes

	courses, err := dao.GetAllCourse(ctx)

	if err != nil {
		return res, err
	}

	for _, val := range courses {

		var course admin_dto.CourseAllRes
		course.ID = val.ID
		course.Title = val.Title
		course.Desc = val.Desc
		course.Author = getTeacherNameByCourseID(ctx, val.ID)

		res = append(res, course)
	}

	return res, nil
}

func getTeacherNameByCourseID(ctx *gin.Context, courseID int) []string {
	teacherID := dao.GetTeacherIdByCourseID(ctx, courseID)
	userName := dao.GetTeacherNameByID(ctx, teacherID)
	return userName
}

// 删除课程，相关表：
func CourseDel(ctx *gin.Context, id int) (int, cerror.Cerror) {

	//删除课程表courese
	courseID, err := dao.DelCourse(ctx, id)

	if err != nil {
		return 0, err
	}

	//删除course_teacher表

	//删除course_student

	return courseID, nil
}

//根据课程id获取课程缩略图图片的url
func getPictureUrlByCourseID(ctx *gin.Context, courseID int) (int, string, string) {

	var url string
	var name string
	pictureID := dao.GetPictureIDByCourseID(ctx, courseID)
	if pictureID == common.FailedID {
		logger.Errorc(ctx, "[%s] getPictureUrlByCourseID fail,pic can not found", "method")
		return pictureID, "", ""
	} else {
		url = dao.GetResourceContentFromID(ctx, pictureID)
		if url == "" {
			logger.Errorc(ctx, "[%s] GetResourceContentFromID fail,url can not found", "method")
		}

		name, _ = dao.GetResourceTitleByID(ctx, pictureID)
	}

	return pictureID, url, name

}

func getVedioUrlByCourseID(ctx *gin.Context, courseID int) (int, string, string) {
	var vedio dto.Vedio
	vedioID := dao.GetVedioIdByCourseID(ctx, courseID)
	vedio.Url, vedio.Desc, vedio.Name = dao.GetVedioInfoFromID(ctx, vedioID)

	return vedioID, vedio.Url, vedio.Name
}

// 查询课程详细信息
func Course(ctx *gin.Context, id int) (admin_dto.CourseRes, cerror.Cerror) {

	var res admin_dto.CourseRes

	course, err := dao.GetCourse(ctx, id)

	if err != nil {
		return res, err
	}

	res.Title = course.Title
	res.Desc = course.Desc
	res.Author = getTeacherNameByCourseID(ctx, course.ID)
	res.PictureID, res.PictureUrl, res.PictureName = getPictureUrlByCourseID(ctx, course.ID)

	res.VedioID, res.VedioUrl, res.VedioName = getVedioUrlByCourseID(ctx, course.ID)

	return res, nil
}

//增加课程
func AddCourse(ctx *gin.Context, param admin_dto.AddCourseParam) (int, cerror.Cerror) {
	//检查输入参数是否合理，检查是否存在课程id

	return dao.AddCourse(ctx, param)

}

func UpdateCourse(ctx *gin.Context, param admin_dto.UpdateCourseParam) (int, cerror.Cerror) {
	//course
	dao.UpdateCourse(ctx, param)

	return 0, nil
}

func AddCourseExtra(ctx *gin.Context, param admin_dto.CourseExtraParam) (admin_dto.AddCourseExtraRes, cerror.Cerror) {
	var res admin_dto.AddCourseExtraRes
	res.CourseResourceID = param.CourseID

	//course_teacher
	var teacherParam admin_dto.CourseTeacherParam
	teacherParam.CourseID = param.CourseID
	teacherParam.TeacherId = param.AuthorID
	res.TeachersResourceID, _ = dao.AddCourseTeacher(ctx, teacherParam)

	//course_resource
	var vedioResourceParam admin_dto.CourseResourceParam
	vedioResourceParam.CourseID = param.CourseID
	vedioResourceParam.ResourceID = param.VedioID
	vedioResourceParam.Type = common.CourseVedio
	vedioResourceParam.Index = 0
	vid, _ := dao.AddCourseResource(ctx, vedioResourceParam)
	res.VedioResourceID = vid

	//course_resource
	var picResourceParam admin_dto.CourseResourceParam
	picResourceParam.CourseID = param.CourseID
	picResourceParam.ResourceID = param.PictureID
	picResourceParam.Type = common.CoursePicture
	picResourceParam.Index = 0
	pid, _ := dao.AddCourseResource(ctx, picResourceParam)

	res.PictureResourceID = pid

	return res, nil
}

func UpdateCourseExtra(ctx *gin.Context, param admin_dto.CourseExtraParam) cerror.Cerror {
	//todo dao.updateTeacher(ctx,param)
	dao.UpdateCourseExtra(ctx, param)

	return nil
}

//根据课程id获取章节目录
func getLessonListByCourseID(ctx *gin.Context, courseID int) ([]admin_dto.CourseLessonAllRes, cerror.Cerror) {

	var lessonList []admin_dto.CourseLessonAllRes

	objList, err1 := dao.GetLessonListByCourseID(ctx, courseID)
	if err1 != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err1.Error())
		return lessonList, cerror.NewCerror(common.Failed, err1.Error())
	}

	for _, val := range objList {
		var item admin_dto.CourseLessonAllRes

		lesson, err := dao.GetLesson(ctx, val.ResourceID)
		if err != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err.Error())
			return lessonList, cerror.NewCerror(common.Failed, err.Error())
		}

		item.ID = val.ResourceID
		item.Title = lesson.Title //dao.GetLessonTitleByID(ctx, val.ResourceID)
		item.Desc = lesson.Desc
		//item.Index = val.Index

		lessonList = append(lessonList, item)
	}

	return lessonList, nil
}

func CourseLessonAll(ctx *gin.Context, id int) ([]admin_dto.CourseLessonAllRes, cerror.Cerror) {

	res, err := getLessonListByCourseID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func UpdateCourseTeacher(ctx *gin.Context, param admin_dto.CourseTeacherParam) (int, cerror.Cerror) {

	return dao.UpdateCourseTeacher(ctx, param)

}

func CourseUpdate(ctx *gin.Context, param admin_dto.UpdateCourse) (int, cerror.Cerror) {
	//更新基本信息
	var param1 admin_dto.UpdateCourseParam
	param1.ID = param.CourseID
	param1.Title = param.Title
	param1.Desc = param.Desc
	UpdateCourse(ctx, param1)
	//更新额外信息
	var param2 admin_dto.CourseExtraParam
	param2.CourseID = param.CourseID
	param2.VedioID = param.VedioID
	param2.PictureID = param.PictureID
	UpdateCourseExtra(ctx, param2)
	//更新教师
	var param3 admin_dto.CourseTeacherParam
	param3.CourseID = param.CourseID
	param3.TeacherId = param.AuthorID
	UpdateCourseTeacher(ctx, param3)

	return param.CourseID, nil
}

func CourseLessonDel(ctx *gin.Context, CourseID int, lessonID int) cerror.Cerror {

	//1 删除章节
	dao.DelLesson(ctx, CourseID, lessonID)

	//2 删除章节与课程关联
	//删除lesson by 课程id和章节index
	var rtype int = common.CourseLesson
	err := dao.DelCourseResource(ctx, CourseID, lessonID, rtype)

	return err

}

package service

import (
	"edu-imp/config"
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//增加课程资源
func AddCourseResource(ctx *gin.Context, param admin_dto.CourseResourceParam) (int, cerror.Cerror) {

	return dao.AddCourseResource(ctx, param)
}

//增加教师-课程信息
func AddCourseTeacher(ctx *gin.Context, param admin_dto.CourseTeacherParam) ([]int, cerror.Cerror) {

	return dao.AddCourseTeacher(ctx, param)
}

//根据课程id获取课程缩略图图片的url
func getPictureUrlByCourseID(ctx *gin.Context, courseID int) string {

	var url string
	pictureID := dao.GetPictureIDByCourseID(ctx, courseID)
	if pictureID == common.FailedID {
		logger.Errorc(ctx, "[%s] getPictureUrlByCourseID fail,pic can not found", "method")
	} else {
		url = dao.GetResourceContentFromID(ctx, pictureID)
		if url == "" {
			logger.Errorc(ctx, "[%s] GetResourceContentFromID fail,url can not found", "method")
		}
	}

	return url

}

//首页获取推荐课程
func GetCourseRecommend(ctx *gin.Context) []dto.CourseRecommendRes {

	var res []dto.CourseRecommendRes

	courseNumberStr := fmt.Sprintf("%d", config.Vipper.Get("Course.CourseNumber"))
	courseNumber, _ := strconv.ParseInt(courseNumberStr, 10, 32)
	course := dao.GetNCourse(ctx, int(courseNumber))

	for _, val := range course {
		var courseRecommend dto.CourseRecommendRes
		courseRecommend.CourseName = val.Title
		courseRecommend.CourseDesc = val.Desc
		courseRecommend.CourseID = val.ID
		courseRecommend.Teacher = getTeacherNameByCourseID(ctx, val.ID)
		courseRecommend.PictureUrl = getPictureUrlByCourseID(ctx, val.ID)

		res = append(res, courseRecommend)
	}

	return res
}

//获取课程列表
func CourseALL(ctx *gin.Context, param dto.CourseAllParam) (dto.CourseAllRes, cerror.Cerror) {

	courseData, err := dao.GetCoursePagedData(ctx, param)
	count, _ := dao.GetCourseCount(ctx)

	//pageSize := fmt.Sprintf("%d", config.Vipper.Get("Course.PageSize"))
	//ps, _ := strconv.ParseInt(pageSize, 10, 64)

	//var page common.Page
	//page = page.CreatePageInfo(param.CurrentPage, ps, count)

	var courseList []dto.CourseRecommendRes
	for _, val := range courseData {
		//fmt.Println(i, val)
		var course dto.CourseRecommendRes
		course.CourseID = val.ID
		course.CourseName = val.Title
		course.CourseDesc = val.Desc
		course.Teacher = getTeacherNameByCourseID(ctx, val.ID)
		course.PictureUrl = getPictureUrlByCourseID(ctx, val.ID)

		courseList = append(courseList, course)

	}

	var courseAllRes dto.CourseAllRes
	courseAllRes.TotalNumber = count
	//courseAllRes.CurrentPage = param.CurrentPage
	//courseAllRes.Pages = page.Nums
	//courseAllRes.PageSize = ps
	courseAllRes.CourseList = courseList

	return courseAllRes, err
}

func getTeacherNameByCourseID(ctx *gin.Context, courseID int) []string {
	teacherID := dao.GetTeacherIdByCourseID(ctx, courseID)
	userName := dao.GetTeacherNameByID(ctx, teacherID)
	return userName
}

//根据课程id获取章节目录
func getLessonListByCourseID(ctx *gin.Context, courseID int) ([]dto.LessonItem, cerror.Cerror) {

	var lessonList []dto.LessonItem

	daolist, err1 := dao.GetLessonListByCourseID(ctx, courseID)
	if err1 != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err1.Error())
		return lessonList, cerror.NewCerror(common.Failed, err1.Error())
	}

	for _, val := range daolist {
		var item dto.LessonItem

		lesson, err := dao.GetLesson(ctx, val.ResourceID)
		if err != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", err.Error())
			return lessonList, cerror.NewCerror(common.Failed, err.Error())
		}

		item.LessonId = val.ResourceID
		item.LessonTitle = lesson.Title //dao.GetLessonTitleByID(ctx, val.ResourceID)
		item.Desc = lesson.Desc
		item.Index = val.Index

		lessonList = append(lessonList, item)
	}

	return lessonList, nil
}

func getVedioByCourseID(ctx *gin.Context, courseID int) dto.Vedio {
	var vedio dto.Vedio
	res := dao.GetVedioIdByCourseID(ctx, courseID)
	vedio.Url, vedio.Desc, vedio.Name = dao.GetVedioInfoFromID(ctx, res)

	return vedio
}

//获取课程的信息，包括主讲人，课程概述，引导视频，课堂章节标题等
func CourseOverview(ctx *gin.Context, param dto.IDParam) (dto.CourseOverviewRes, cerror.Cerror) {
	var res dto.CourseOverviewRes

	course := dao.GetCourseByID(ctx, param.ID) //课程的简要信息
	res.Introduce = course.Desc
	res.Title = course.Title

	res.Teacher = getTeacherNameByCourseID(ctx, course.ID)     //课程任课教师
	res.PictureUrl = getPictureUrlByCourseID(ctx, course.ID)   //课程图片
	res.GuideVedio = getVedioByCourseID(ctx, course.ID)        //课程介绍视频
	lessonList, err := getLessonListByCourseID(ctx, course.ID) //课程章节

	res.LessonList = lessonList

	return res, err
}

//func GetCourseIDByLessonID()  {
//
//}

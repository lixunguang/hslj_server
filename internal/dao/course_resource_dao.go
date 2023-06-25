package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseResource struct {
	BaseModel
	CourseID   int            `gorm:"column:course_id"`
	ResourceID int            `gorm:"column:resource_id"`
	Type       int            `gorm:"column:type"`  //资源用途类型，展示图片，引导视频，章节id 资源用途：1展示图片，2 引导视频，3 章节id
	Index      int            `gorm:"column:index"` //排序
	DeletedAt  gorm.DeletedAt //支持软删
}

func (CourseResource) TableName() string {
	return "course_resource"
}

//增加课程资源
func AddCourseResource(ctx *gin.Context, param admin_dto.CourseResourceParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	courseResource := CourseResource{CourseID: param.CourseID, ResourceID: param.ResourceID,
		Type: param.Type, Index: param.Index}

	result := mysqlDB.Create(&courseResource)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return courseResource.ID, nil
}

//删除记录，软删除
func DelCourseResource(ctx *gin.Context, courseID int, resourceID int, rtype int) cerror.Cerror {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Where("course_id = ? and resource_id = ? and `type` = ?", courseID, resourceID, rtype).Delete(&CourseResource{})
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return nil
}

//根据课程id获取课程缩略图图片的url
func GetPictureIDByCourseID(ctx *gin.Context, courseID int) int {
	mysqlDB := mysql.GetDB()

	var courseResource CourseResource

	result := mysqlDB.Where("type = ? and course_id= ? ", common.CoursePicture, courseID).First(&courseResource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return 0
	}

	return courseResource.ResourceID
}

func GetLessonListByCourseID(ctx *gin.Context, courseID int) ([]CourseResource, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var lessonList []CourseResource

	result := mysqlDB.Where("type = ? and course_id = ?", common.CourseLesson, courseID).Find(&lessonList)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return lessonList, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonList, nil

}

func GetVedioIdByCourseID(ctx *gin.Context, courseID int) int {

	mysqlDB := mysql.GetDB()

	var res int = 0
	var vedio CourseResource

	result := mysqlDB.Where("type = ? and course_id = ?", common.CourseVedio, courseID).First(&vedio)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res
	}

	res = vedio.ResourceID
	return res

}

//更新视频，图片，章节信息
func UpdateCourseExtra(ctx *gin.Context, param admin_dto.CourseExtraParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var cr CourseResource

	//更新图片
	result1 := mysqlDB.Model(&cr).Where("course_id = ? and type = 1", param.CourseID).Updates(CourseResource{ResourceID: param.PictureID})
	if result1.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result1.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result1.Error.Error())
	}
	//更新视频
	result2 := mysqlDB.Model(&cr).Where("course_id = ? and type = 2", param.CourseID).Updates(CourseResource{ResourceID: param.VedioID})
	if result2.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result2.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result2.Error.Error())
	}

	//更新章节信息
	/*result3 := mysqlDB.Model(&cr).Where("id = ? and type = 3", param.CourseID).Updates(CourseResource{ResourceID:param.VedioID})
	if result3.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result3.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result3.Error.Error())
	}
	*/
	return param.CourseID, nil
}

package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type CourseTeacher struct {
	BaseModel
	CourseID  int `gorm:column:course_id`
	TeacherID int `gorm:column:teacher_id`
}

func (CourseTeacher) TableName() string {
	return "course_teacher"
}

//增加课程资源
//此函数最早支持增加多个课程的教师，后来改为了一个
func AddCourseTeacher(ctx *gin.Context, param admin_dto.CourseTeacherParam) ([]int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res []int
	//for _, val := range param.TeacherIds {
	courseTeacher := CourseTeacher{CourseID: param.CourseID, TeacherID: param.TeacherId}
	result := mysqlDB.Create(&courseTeacher)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.FailedID, result.Error.Error())
	}
	res = append(res, courseTeacher.ID)
	//}

	return res, nil
}

//根据courseid 获取teacherid
func GetTeacherIdByCourseID(ctx *gin.Context, courseID int) []int {
	mysqlDB := mysql.GetDB()

	var res []int
	var courseTeachers []CourseTeacher

	result := mysqlDB.Where("course_id = ? ", courseID).Find(&courseTeachers) //有多个的话
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, courseID)
		return res
	}

	for _, val := range courseTeachers {
		res = append(res, val.TeacherID)
	}

	return res
}

func UpdateCourseTeacher(ctx *gin.Context, param admin_dto.CourseTeacherParam) (int, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	//1：del
	var ct CourseTeacher
	result1 := mysqlDB.Where("course_id = ?", param.CourseID).Delete(&ct)
	if result1.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result1.Error)
		return param.CourseID, cerror.NewCerror(common.Failed, result1.Error.Error())
	}

	//2：add
	AddCourseTeacher(ctx, param)

	return param.CourseID, nil

}

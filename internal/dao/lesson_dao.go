package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Lesson struct {
	BaseModel
	Title     string         `gorm:"column:title"`
	Desc      string         `gorm:"column:desc"`
	DeletedAt gorm.DeletedAt //支持软删
}

func (Lesson) TableName() string {
	return "lesson"
}

//增加课程
func AddLesson(ctx *gin.Context, param dto.Lesson) (dto.IDRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	lesson := Lesson{Title: param.Title, Desc: param.Desc}

	result := mysqlDB.Create(&lesson)
	var res dto.IDRes

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	res.ID = lesson.ID
	return res, nil
}

//删除课程
func DelLesson(ctx *gin.Context, CourseID int, lessonID int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Delete(&Lesson{}, lessonID)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return 0, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonID, nil
}

func GetLesson(ctx *gin.Context, id int) (Lesson, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var lesson Lesson
	result := mysqlDB.Where("id = ? ", id).First(&lesson)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return lesson, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lesson, nil
}

func LessonUpdate(ctx *gin.Context, CourseID int, lessonID int, title string, desc string) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Where("id = ? ", lessonID).Updates(Lesson{Title: title, Desc: desc})
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return lessonID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonID, nil
}

func GetLessonTitleByID(ctx *gin.Context, id int) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var lesson Lesson
	result := mysqlDB.Where("id = ? ", id).First(&lesson)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return "", cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lesson.Title, nil
}

//-------------------------------------------------------------后台

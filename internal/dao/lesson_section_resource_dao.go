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

type LessonSectionResource struct {
	BaseModel
	LessonSectionID int `gorm:"column:lesson_section_id"` //lesson_section id
	ResourceID      int `gorm:"column:resource_id"`       //资源id
	Index           int `gorm:"column:index"`             //排序

	DeletedAt gorm.DeletedAt //软删
}

func (LessonSectionResource) TableName() string {
	return "lesson_section_resource"
}

func GetLessonSectionResourceBySectionIDType(ctx *gin.Context, lessonSectionID int) ([]LessonSectionResource, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var lessonSectionResource []LessonSectionResource
	result := mysqlDB.Where("lesson_section_id = ?", lessonSectionID).Find(&lessonSectionResource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonSectionResource, nil
}

func AddLessonSectionResource(ctx *gin.Context, param dto.LessonSectionResource) (dto.IDRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	lessonResource := LessonSectionResource{LessonSectionID: param.SectionID, ResourceID: param.ResourceID, Index: param.Index}

	result := mysqlDB.Create(&lessonResource)

	var res dto.IDRes
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	res.ID = lessonResource.ID
	return res, nil
}

func DelLessonSectionResource(ctx *gin.Context, param dto.LessonSectionDelResource) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var lsr LessonSectionResource
	result := mysqlDB.Where("lesson_section_id = ? and resource_id = ?", param.SectionID, param.ResourceID).Delete(&lsr)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return int(result.RowsAffected), cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return int(result.RowsAffected), nil
}

package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

//课堂章节的大目录，如教学内容，教学参考，实验资源，作业，只记录到目录层次
type LessonSection struct {
	ID       int    `gorm:"primary_key"`
	LessonID int    `gorm:"lesson_id"`
	Title    string `gorm:"title"`
	Index    int    `gorm:"index"`
}

func (LessonSection) TableName() string {
	return "lesson_section"
}

func AddLessonSection(ctx *gin.Context, param dto.LessonSection) (int, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	lessonSection := LessonSection{LessonID: param.LessonID, Title: param.Title, Index: param.Index}
	result := mysqlDB.Create(&lessonSection)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonSection.ID, nil
}

//根据lessonid 和type 获取lesson_section 的id,返回值为一个id
func GetLessonSectionIDByLessonIDType(ctx *gin.Context, lessonID int, sectionIndex int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var lessonSection LessonSection
	result := mysqlDB.Where("lesson_id = ? and `index` = ? ", lessonID, sectionIndex).First(&lessonSection)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, lessonID=%d", result.Error, lessonID)
		return lessonSection.ID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return lessonSection.ID, nil
}

func GetLessonSectionByLessonID(ctx *gin.Context, lessonID int) ([]dto.LessonSection, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var res []dto.LessonSection

	var lessonSection []LessonSection
	result := mysqlDB.Where("lesson_id=?", lessonID).Find(&lessonSection)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, lessonID=%d", result.Error, lessonID)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	for _, val := range lessonSection {
		var item dto.LessonSection
		item.Title = val.Title
		item.Index = val.Index
		item.LessonID = val.LessonID

		res = append(res, item)
	}

	return res, nil
}

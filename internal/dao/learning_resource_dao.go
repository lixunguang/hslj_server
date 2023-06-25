package dao

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type LearningResource struct {
	BaseModel
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	LearningID int    `gorm:"column:learning_id"`
	ResourceID int    `gorm:"column:resource_id"`
	Index      int    `gorm:"column:index"`
}

func (LearningResource) TableName() string {
	return "learning_resource"
}

func AddLearningResource(ctx *gin.Context, title string, desc string, param admin_dto.AddLearningResourceParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learningItem = LearningResource{Title: title, Desc: desc, LearningID: param.LearningID, ResourceID: param.ResourceID, Index: param.Index}
	result := mysqlDB.Create(&learningItem)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return learningItem.ID, nil

}

func DelLearningResource(ctx *gin.Context, param admin_dto.DelLearningResourceParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	result := mysqlDB.Where("resource_id = ? and learning_id = ?", param.ResourceID, param.LearningID).Delete(&LearningResource{})
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return param.LearningID, nil

}

func GetLearningResourceList(ctx *gin.Context, learningID int) ([]LearningResource, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learningItemList []LearningResource

	result := mysqlDB.Where("learning_id=?", learningID).Find(&learningItemList)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return learningItemList, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return learningItemList, nil
}

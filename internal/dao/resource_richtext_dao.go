package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

/* 富文本单独出来一个表*/
type ResourceRichText struct {
	BaseModel
	Content string `gorm:"column:content"`
}

func (ResourceRichText) TableName() string {
	return "resource_richtext"
}

//增加富文本资源
func AddResourceRichText(ctx *gin.Context, param dto.ResourceRichText) (int, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	//add
	resource := ResourceRichText{Content: param.Content}

	result := mysqlDB.Create(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=%+v", result.Error)
		return common.Failed, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resource.ID, nil
}

//删除提交的富文本资源
func DelResourceRichtext(ctx *gin.Context, id int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res int //删除的资源id
	var richtext ResourceRichText
	richtext.ID = id

	result := mysqlDB.Delete(&richtext)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)

		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return res, nil
}

func UpdateResourceRichtext(ctx *gin.Context, param dto.UpdateRichtext) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var richtext ResourceRichText
	richtext.ID = param.ID

	result := mysqlDB.Model(&richtext).Where("id = ?", param.ID).Updates(ResourceRichText{Content: param.Content})

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return param.ID, nil
}

func GetResourceRichText(ctx *gin.Context, param dto.IDParam) (dto.ResourceRichText, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res dto.ResourceRichText

	var richText ResourceRichText
	result := mysqlDB.Where("id = ?", param.ID).First(&richText)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 1,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	res.Content = richText.Content
	return res, nil

}

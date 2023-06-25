package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

/* 资源表，富文本单独出来一个表*/

type Resource struct {
	BaseModel
	Title   string `gorm:"column:title"`
	Desc    string `gorm:"column:desc"`
	Type    int8   `gorm:"column:type"` // 在common define里面有定义：1图片，	2 word文档，	3 pdf文档，	4 apps文件，	5 ibe文件，	6文本（未用）	7 富文本。	8 视频
	Content string `gorm:"column:content"`

	DeletedAt gorm.DeletedAt //支持软删
}

func (Resource) TableName() string {
	return "resource"
}

//增加资源
func AddResource(ctx *gin.Context, param dto.Resource) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res int = 0
	resource := Resource{Title: param.Title, Desc: param.Desc, Type: param.Type, Content: param.Content}
	result := mysqlDB.Create(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return resource.ID, nil

}

//删除资源
func DelResource(ctx *gin.Context, IDs []int) ([]int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var resource Resource
	result := mysqlDB.Delete(&resource, IDs)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return []int{}, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return IDs, nil
}

func UpdateResource(ctx *gin.Context, param dto.UpdateResourceParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var resource Resource
	resource.ID = param.ID

	result := mysqlDB.Model(&resource).Where("id = ?", param.ID).Updates(Resource{Title: param.Title, Desc: param.Desc, Type: param.Type})

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return param.ID, nil
}

//根据 id 获取resource
func GetResourceByID(ctx *gin.Context, id int) (Resource, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return resource, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resource, nil
}

//根据 id 获取resource type
func GetResourceTypeByID(ctx *gin.Context, id int) (int8, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return resource.Type, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resource.Type, nil
}

//根据 id 获取resource title
func GetResourceTitleByID(ctx *gin.Context, id int) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return resource.Title, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resource.Title, nil
}

//根据 id 获取desc
func GetResourceDescByID(ctx *gin.Context, id int) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return resource.Title, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return resource.Desc, nil
}

//根据resource id 获取resouce content，注意：当为富文本时，从富文本表里面获取
func GetResourceContentFromID(ctx *gin.Context, id int) string {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return ""
	}

	var content string

	if resource.Type == common.RichType {

		index, _ := strconv.ParseUint(resource.Content, 10, 64) //todo: error
		rr, _ := GetResourceRichText(ctx, dto.IDParam{ID: int(index)})
		content = rr.Content
	} else {
		content = common.GetCDNAddr() + resource.Content
	}

	return content
}

func GetFilePathByID(ctx *gin.Context, id int) (string, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var resource Resource
	result := mysqlDB.Where("id = ? ", id).Find(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return resource.Content, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return resource.Content, nil
}

func GetRichtextIdByID(ctx *gin.Context, id int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? and type= 7 ", id).Find(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return 0, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	res, err := strconv.Atoi(resource.Content)
	if err != nil {
		return 0, cerror.NewCerror(common.FailedID, err.Error())
	}

	return res, nil
}

//根据vedio id 获取vedio url
func GetVedioInfoFromID(ctx *gin.Context, id int) (string, string, string) {
	mysqlDB := mysql.GetDB()

	resource := Resource{}
	result := mysqlDB.Where("id = ? ", id).First(&resource)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return "", "", ""
	}

	url := common.GetCDNAddr() + resource.Content

	return url, resource.Desc, resource.Title
}

package dao

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type Action struct {
	BaseModel
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (Action) TableName() string {
	return "action"
}

//获取所有动作
func GetActionAll(ctx *gin.Context, number int) ([]Action, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var actions []Action
	result := mysqlDB.Order("updated_at desc").Limit(number).Find(&actions)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetActionAll] fail,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return actions, nil
}

// 获取动作详情
func GetActionDetail(ctx *gin.Context, id int) (Action, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	res := Action{}
	result := mysqlDB.Where("id = ? ", id).First(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, id=%d", result.Error, id)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return res, nil
}

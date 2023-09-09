package dao

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"time"
)

type Today struct {
	BaseModel
	LunarDesc  string `gorm:"column:lunar_desc"`
	Txt        string `gorm:"column:txt"`
	PictureUrl string `gorm:"column:picture_url"`
}

func (Today) TableName() string {
	return "today"
}

func GetToday(ctx *gin.Context) (Today, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	timeStr := time.Now().Format("20060102")
	var today Today
	result := mysqlDB.Where("date= ?", timeStr).First(&today)

	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetToday] fail,err=%+v", result.Error)
		return today, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return today, nil
}

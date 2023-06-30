package dao

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dto"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type Location struct {
	BaseModel

	Name     string `gorm:"column:name"`
	Desc     string `gorm:"column:desc"`
	Location string `gorm:"column:location"`
	Rating   int    `gorm:"column:rating"`
}

func (Location) TableName() string {
	return "location"
}

func GetLocation(ctx *gin.Context, param dto.GetLocationListParam) ([]Location, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var location []Location

	result := mysqlDB.Find(&location)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return location, cerror.ErrorDataGet
	}

	return location, nil
}

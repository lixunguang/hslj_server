package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hslj/internal/common"

	"hslj/internal/dto"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type Location struct {
	BaseModel

	Name      string        `gorm:"column:name"`
	Desc      string        `gorm:"column:desc"`
	Loc       string 		`gorm:"column:location"`
	Frequency int           `gorm:"column:frequency"`  //频率
	TimeStr   string        `gorm:"column:time_str"`   //举行时间
	PeopleNum int           `gorm:"column:people_num"` //参与人数
	Rating    int           `gorm:"column:rating"`     //综合评分，包括参与人数，场地，技术水平，文化氛围等
	IsAuth    int           `gorm:"column:is_auth"`
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

func AddLocation(ctx *gin.Context, param dto.AddLocationParam) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var location Location
	location.Desc = param.Desc
	location.Name = param.Name
	location.Frequency = param.Frequency
	location.PeopleNum = param.PeopleNum
	location.TimeStr = param.TimeStr
	location.IsAuth = param.ISAuth
	location.Rating = param.Rating
	location.Loc = "sda"


	var sqlStr string  = "INSERT INTO `location` (`name`,`desc`,`location`,`frequency`,`time_str`,`people_num`,`rating`) VALUES ( '%s', '%s',ST_GeomFromText('POINT(121.474103 31.232862)'),'%d','%s','%d', '%d');"
	sqlStr2 := fmt.Sprintf(sqlStr,param.Name,param.Desc,param.Frequency,param.TimeStr,param.PeopleNum,param.Rating)
	result := mysqlDB.Exec(sqlStr2)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.InvalidID, cerror.ErrorDataAdd
	}

	return location.ID, nil

}

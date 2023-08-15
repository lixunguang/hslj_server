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

	Name      string `gorm:"column:name"`
	Desc      string `gorm:"column:desc"` //主要介绍自然景观，毽子场地，路线，时间，注意事项等。
	Loc       string `gorm:"column:location"`
	Frequency int    `gorm:"column:frequency"`  //频率
	TimeStr   string `gorm:"column:time_str"`   //举行时间
	PeopleNum int    `gorm:"column:people_num"` //参与人数
	Rating    int    `gorm:"column:rating"`     //综合评分，包括参与人数，场地，技术水平，文化氛围等，比如每一个月启动一次评价
	Hot       int    `gorm:"column:hot"`        //热度，一段时间内参与的人数,是一个动态调整的概念，比如每一个月启动一次评价，进入热门场地
	IsAuth    int    `gorm:"column:is_auth"`
}

type LocationRes struct {
	BaseModel

	Name      string `gorm:"column:name"`
	Desc      string `gorm:"column:desc"` //主要介绍自然景观，毽子场地，路线，时间，注意事项等。
	Loc       string `gorm:"column:st_ast"`
	Frequency int    `gorm:"column:frequency"`  //频率
	TimeStr   string `gorm:"column:time_str"`   //举行时间
	PeopleNum int    `gorm:"column:people_num"` //参与人数
	Rating    int    `gorm:"column:rating"`     //综合评分，包括参与人数，场地，技术水平，文化氛围等，比如每一个月启动一次评价
	Hot       int    `gorm:"column:hot"`        //热度，一段时间内参与的人数,是一个动态调整的概念，比如每一个月启动一次评价，进入热门场地
	IsAuth    int    `gorm:"column:is_auth"`

}

func (Location) TableName() string {
	return "location"
}

func GetLocationList(ctx *gin.Context, areaCode int) ([]Location, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var location []Location

	result := mysqlDB.Where("area_code = ?", areaCode).Find(&location)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return location, cerror.ErrorDataGet
	}

	return location, nil
}

func GetLocationHot(ctx *gin.Context) ([]Location, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var location []Location

	result := mysqlDB.Order("hot desc").Limit(5).Find(&location)

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

	var sqlStr string = "INSERT INTO `location` (`name`,`desc`,`location`,`frequency`,`time_str`,`people_num`,`rating`) VALUES ( '%s', '%s',ST_GeomFromText('POINT(121.474103 31.232862)'),'%d','%s','%d', '%d');"
	sqlStr2 := fmt.Sprintf(sqlStr, param.Name, param.Desc, param.Frequency, param.TimeStr, param.PeopleNum, param.Rating)
	result := mysqlDB.Exec(sqlStr2)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.InvalidID, cerror.ErrorDataAdd
	}

	return location.ID, nil

}

func GetLocationDetail(ctx *gin.Context,id int) (LocationRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var location LocationRes

	var sqlStr string = "SELECT `name`,`desc`,st_astext(location) as st_ast,`frequency`,`time_str`,`people_num`,`hot`,`rating`,`is_auth` FROM `location` WHERE id= %d;"
	sqlStr2 := fmt.Sprintf(sqlStr, id)
	result := mysqlDB.Raw(sqlStr2).Find(&location)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return location, cerror.ErrorDataGet
	}

	return location, nil
}

package dao

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/dto"
	"hslj/internal/model/mysql"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
)

type Record struct {
	BaseModel
	UserLoginID string `gorm:"column:user_login_id"`
	LocationID  int    `gorm:"column:location_id"`
}

func (Record) TableName() string {
	return "record"
}

func AddRecord(ctx *gin.Context, param dto.AddRecordParam) (int, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	record := Record{UserLoginID: param.UserLoginID, LocationID: param.LocationID}
	result := mysqlDB.Create(&record)

	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.AddRecord] fail,err=%+v", result.Error)
		return common.InvalidID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return record.ID, nil
}

func GetRecordByUserID(ctx *gin.Context, param dto.IDParam) ([]Record, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var records []Record
	result := mysqlDB.Order("updated_at desc").Where("user_login_id = ?", param.ID).Find(&records)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetRecordByUserID] fail,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return records, nil
}

func GetRecordByLocationID(ctx *gin.Context, param dto.IDParam) ([]Record, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var records []Record
	result := mysqlDB.Order("updated_at desc").Where("location_id = ?", param.ID).Find(&records)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsdao.GetRecordByLocationID] fail,err=%+v", result.Error)
		return nil, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return records, nil
}

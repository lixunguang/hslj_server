package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type FileTable struct {
	BaseModel
	OriginalName string `gorm:"column:original_name"`
	StorageName  string `gorm:"column:storage_name"`
}

func (FileTable) TableName() string {
	return "file"
}

func GetOriginalNameFromStorageName(ctx *gin.Context, originalName string) string {
	mysqlDB := mysql.GetDB()

	var fileTable FileTable
	result := mysqlDB.Where("OriginalName=?", originalName).First(&fileTable)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return fileTable.OriginalName
	}

	return fileTable.OriginalName
}

func AddFileTable(ctx *gin.Context, param dto.FileTable) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	fileTable := FileTable{OriginalName: param.OriginalName, StorageName: param.StorageName}
	result := mysqlDB.Create(&fileTable)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return fileTable.ID, nil
}

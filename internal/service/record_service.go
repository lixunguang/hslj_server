package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
	"hslj/pkg/util"
)

func AddRecord(ctx *gin.Context, param dto.AddRecordParam) (int, cerror.Cerror) {

	id, err := dao.AddRecord(ctx, param)

	return id, err
}

func GetRecordByUserID(ctx *gin.Context, param dto.IDParam) ([]dto.LocationRecordRes, cerror.Cerror) {

	var res []dto.LocationRecordRes
	var locationItem dto.LocationRecordRes

	records, err := dao.GetRecordByUserID(ctx, param)
	for _, val := range records {

		location, err2 := dao.GetLocationDetail(ctx, val.LocationID)
		fmt.Println(err2)


		locationItem.Name = location.Name

		locationItem.DataStr = val.UpdatedAt.Format(util.FormatDate)

		res = append(res, locationItem)
	}

	return res, err
}

func GetRecordByLocationID(ctx *gin.Context, param dto.IDParam) ([]dto.UserRecordRes, cerror.Cerror) {

	var res []dto.UserRecordRes
	var userItem dto.UserRecordRes

	records, err := dao.GetRecordByLocationID(ctx, param)
	for _, val := range records {

		location, err2 := dao.GetUser(ctx, val.UserLoginID)
		fmt.Println(err2)

		userItem.LoginID = location.LoginID
		userItem.Name = location.Name
		userItem.DataStr = val.UpdatedAt.Format(util.FormatDate)

		res = append(res, userItem)
	}

	return res, err
}

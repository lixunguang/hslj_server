package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func AddRecord(ctx *gin.Context, param dto.AddRecordParam) (int, cerror.Cerror) {

	id, err := dao.AddRecord(ctx, param)

	return id, err
}

func GetRecordByUserID(ctx *gin.Context, param dto.IDParam) ([]dto.LocationRes, cerror.Cerror) {

	var res []dto.LocationRes
	var locationItem dto.LocationRes

	records, err := dao.GetRecordByUserID(ctx, param)
	for _, val := range records {

		location, err2 := dao.GetLocationDetail(ctx, val.LocationID)
		fmt.Println(err2)

		locationItem.Desc = location.Desc
		locationItem.Name = location.Name

		res = append(res, locationItem)
	}

	return res, err
}

func GetRecordByLocationID(ctx *gin.Context, param dto.IDParam) ([]dto.LocationRes, cerror.Cerror) {

	var res []dto.LocationRes
	var locationItem dto.LocationRes

	records, err := dao.GetRecordByLocationID(ctx, param)
	for _, val := range records {

		location, err2 := dao.GetLocationDetail(ctx, val.LocationID)
		fmt.Println(err2)

		locationItem.Desc = location.Desc
		locationItem.Name = location.Name

		res = append(res, locationItem)
	}

	return res, err
}

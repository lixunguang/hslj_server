package service

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func GetLocation(ctx *gin.Context, param dto.GetLocationListParam) ([]dto.GetLocationListRes, cerror.Cerror) {

	var res []dto.GetLocationListRes

	location, err := dao.GetLocation(ctx, param)
	for _, val := range location {
		var item dto.GetLocationListRes
		item.ID = val.ID
		item.Rating = val.Rating
		item.Name = val.Name
		item.Desc = val.Desc
		item.Longitude = 12
		item.Latitude = 34
		res = append(res, item)

	}

	return res, err
}

func GetLocationHot(ctx *gin.Context) ([]dto.GetLocationListRes, cerror.Cerror) {

	var res []dto.GetLocationListRes

	location, err := dao.GetLocationHot(ctx)
	for _, val := range location {
		var item dto.GetLocationListRes
		item.ID = val.ID
		item.Rating = val.Rating
		item.Name = val.Name
		item.Desc = val.Desc
		item.Longitude = 12
		item.Latitude = 34
		res = append(res, item)

	}

	return res, err
}
func AddLocation(ctx *gin.Context, param dto.AddLocationParam) (int, cerror.Cerror) {

	id, err := dao.AddLocation(ctx, param)

	return id, err
}

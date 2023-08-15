package service

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/common"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func GetLocation(ctx *gin.Context, id int) (dto.LocationRes, cerror.Cerror) {

	var res dto.LocationRes

	location, err := dao.GetLocationDetail(ctx, id)
	res.ID = location.ID
	res.Desc = location.Desc
	res.Name = location.Name
	res.Rating = location.Rating
	res.Latitude =12
	res.Longitude = 34

	return res, err
}

func GetLocationList(ctx *gin.Context, param dto.GetLocationListParam) ([]dto.LocationRes, cerror.Cerror) {
	var res []dto.LocationRes

	location, err := dao.GetLocationList(ctx, param.LocationCode)
	for _, val := range location {
		var item dto.LocationRes
		item.ID = val.ID
		item.Rating = val.Rating
		item.Name = common.GetShortStr(val.Name,common.TitleLength)
		item.Desc = common.GetShortStr(val.Desc,common.DescLength)
		item.Longitude = 12
		item.Latitude = 34
		res = append(res, item)
	}

	return res, err
}

func GetLocationHot(ctx *gin.Context) ([]dto.LocationRes, cerror.Cerror) {

	var res []dto.LocationRes

	location, err := dao.GetLocationHot(ctx)
	for _, val := range location {
		var item dto.LocationRes
		item.ID = val.ID
		item.Name = common.GetShortStr(val.Name,common.TitleLength)
		item.Desc = common.GetShortStr(val.Desc,common.DescLength)
		item.Rating = val.Rating
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

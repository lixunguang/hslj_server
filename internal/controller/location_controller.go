package controller

import (
	"github.com/gin-gonic/gin"
	"hslj/pkg/util"
)

type LocationRes struct {
	Name      string
	Desc      string
	Latitude  float64
	Longitude float64
	Rating    int
	//OpenTime time.Time
	//CloseTime time.Time
	//UserReviews []int

}

func GetLocationList(ctx *gin.Context) {
	var res []LocationRes

	var item LocationRes
	item.Name = "北京颐和园"
	item.Desc = "颐和园描述。。。"
	item.Latitude = 1.11
	item.Longitude = 2.22
	item.Rating = 4
	//item.OpenTime =

	res = append(res, item)

	item.Name = "北京天坛"
	item.Desc = "天坛描述。。。"
	item.Latitude = 3.3333
	item.Longitude = 4.4444
	item.Rating = 4
	res = append(res, item)

	util.SuccessJson(ctx, res)
}

func AddLocation(ctx *gin.Context) {

	res := "addok!"

	util.SuccessJson(ctx, res)
}

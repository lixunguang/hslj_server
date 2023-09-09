package service

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func GetToday(ctx *gin.Context) (dto.Today, cerror.Cerror) {

	var res dto.Today
	today, err := dao.GetToday(ctx)

	res.LunarDesc = today.LunarDesc
	res.Txt = today.Txt
	res.PictureUrl = today.PictureUrl

	return res, err
}

func GetTodayShort(ctx *gin.Context) (dto.TodayShort, cerror.Cerror) {

	var res dto.TodayShort
	today, err := dao.GetToday(ctx)

	if len(today.LunarDesc) > 0 {
		res.Desc = today.LunarDesc
	} else {
		res.Desc = today.Txt
	}

	return res, err
}

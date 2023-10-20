package service

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func ActionAll(ctx *gin.Context) ([]dto.ActionAllRes, cerror.Cerror) {

	var res []dto.ActionAllRes

	actions, err := dao.GetActionAll(ctx, 1000)

	for _, val := range actions {
		var item dto.ActionAllRes
		item.ID = val.ID
		item.Title = val.Title

		res = append(res, item)
	}

	return res, err
}

func ActionDetail(ctx *gin.Context, id dto.IDParam) (dto.ActionDetailRes, cerror.Cerror) {
	var res dto.ActionDetailRes

	action, err := dao.GetActionDetail(ctx, id.ID)

	res.ID = action.ID
	res.Title = action.Title
	res.Content = action.Content

	return res, err
}

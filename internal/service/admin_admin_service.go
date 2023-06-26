package service

import (
	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

func AddAdmin(ctx *gin.Context, param dto.Admin) (dto.AdminRes, cerror.Cerror) {

	var getParam dto.AdminParam
	getParam.Name = param.Name
	adminArray, _ := dao.GetAdmin(ctx, getParam)

	if len(adminArray) == 0 {
		return dao.AddAdmin(ctx, param)
	}

	var res dto.AdminRes
	return res, cerror.ErrorUserExist
}

func DelAdmin(ctx *gin.Context, param dto.AdminParam) ([]dto.AdminRes, cerror.Cerror) {

	var res []dto.AdminRes
	admins, err := dao.DelAdmin(ctx, param)

	for _, v := range admins {
		var t dto.AdminRes
		t.Name = v.Name
		t.ID = v.ID

		res = append(res, t)
	}

	return res, err
}

func GetAdmin(ctx *gin.Context, param dto.AdminParam) ([]dao.Admin, cerror.Cerror) {
	return dao.GetAdmin(ctx, param)
}

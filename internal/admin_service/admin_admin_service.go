package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dao"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddAdmin(ctx *gin.Context, param admin_dto.Admin) (admin_dto.AdminRes, cerror.Cerror) {

	var getParam admin_dto.AdminParam
	getParam.Name = param.Name
	adminArray, _ := dao.GetAdmin(ctx, getParam)

	if len(adminArray) == 0 {
		return dao.AddAdmin(ctx, param)
	}

	var res admin_dto.AdminRes
	return res, cerror.ErrorUserExist
}

func DelAdmin(ctx *gin.Context, param admin_dto.AdminParam) ([]admin_dto.AdminRes, cerror.Cerror) {

	var res []admin_dto.AdminRes
	admins, err := dao.DelAdmin(ctx, param)

	for _, v := range admins {
		var t admin_dto.AdminRes
		t.Name = v.Name
		t.ID = v.ID

		res = append(res, t)
	}

	return res, err
}

func GetAdmin(ctx *gin.Context, param admin_dto.AdminParam) ([]dao.Admin, cerror.Cerror) {
	return dao.GetAdmin(ctx, param)
}

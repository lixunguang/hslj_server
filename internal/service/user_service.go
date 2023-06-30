package service

import (

	"github.com/gin-gonic/gin"
	"hslj/internal/dao"
	"hslj/internal/dto"
	"hslj/pkg/cerror"
)

// 增加用户
func AddUser(ctx *gin.Context, param dto.User) (dto.AddUserRes, cerror.Cerror) {

	var res dto.AddUserRes

	getResult, _ := dao.GetUser(ctx, param.LoginID)
	if getResult.LoginID == "" { //未查到
		createRes, err := dao.CreateUser(ctx, param)
		if err == nil {
			return dto.AddUserRes{ID: createRes.ID, Name: createRes.Name}, nil
		}

		return res, cerror.ErrorDataAdd
	}

	return res, cerror.ErrorUserExist

}

func DelUser(ctx *gin.Context, login_id string) (string, cerror.Cerror) {

	return dao.DelUser(ctx, login_id)
}

func UpdateUser(ctx *gin.Context, param dto.User) (string, cerror.Cerror) {

	return dao.UpdateUser(ctx, param)
}

func AllUser(ctx *gin.Context) ([]dto.User, cerror.Cerror) {
	var res []dto.User

	userRes, err := dao.AllUser(ctx)
	if err != nil {
		return res, err
	}

	for _, val := range userRes {
		var item dto.User
		item.Name = val.Name
		item.LoginID = val.LoginID
		item.OrganizationID = val.OrganizationID
		item.OrganizationID = val.OrganizationID
		item.Password = val.Password
		res = append(res, item)
	}

	return res, nil

}

func GetUser(ctx *gin.Context, loginID string) (dto.UserRes, cerror.Cerror) {
	return dao.GetUser(ctx, loginID)
}

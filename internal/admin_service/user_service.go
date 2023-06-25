package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dao"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

// 增加用户
func AddUser(ctx *gin.Context, param admin_dto.User) (admin_dto.AddUserRes, cerror.Cerror) {

	var res admin_dto.AddUserRes

	getResult, _ := dao.GetUser(ctx, param.LoginID)
	if getResult.LoginID == "" { //未查到
		createRes, err := dao.CreateUser(ctx, param)
		if err == nil {
			return admin_dto.AddUserRes{ID: createRes.ID, Name: createRes.Name}, nil
		}

		return res, cerror.ErrorDataAdd
	}

	return res, cerror.ErrorUserExist

}

func DelUser(ctx *gin.Context, login_id string) (string, cerror.Cerror) {

	return dao.DelUser(ctx, login_id)
}

func UpdateUser(ctx *gin.Context, param admin_dto.User) (string, cerror.Cerror) {

	return dao.UpdateUser(ctx, param)
}

func AllUser(ctx *gin.Context) ([]admin_dto.User, cerror.Cerror) {
	var res []admin_dto.User

	userRes, err := dao.AllUser(ctx)
	if err != nil {
		return res, err
	}

	for _, val := range userRes {
		var item admin_dto.User
		item.Name = val.Name
		item.LoginID = val.LoginID
		item.OrganizationID = val.OrganizationID
		item.OrganizationID = val.OrganizationID
		item.Password = val.Password
		res = append(res, item)
	}

	return res, nil

}

func GetUser(ctx *gin.Context, loginID string) (admin_dto.UserRes, cerror.Cerror) {
	return dao.GetUser(ctx, loginID)
}

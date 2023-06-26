package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"hslj/internal/common"
	"hslj/internal/dto"
	"hslj/internal/middleware"
	"hslj/internal/service"
	"hslj/pkg/cerror"
	"hslj/pkg/logger"
	"hslj/pkg/util"
)

// 权限相关接口

// gwt，token串
// token存储，压缩，内存，数据库
// 重复登录，存储新的token，之前token失效。
// 过期检查

//检查token是否过期
func CheckAdminAuth(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	logger.Infoc(ctx, "[CheckAuth Controller] Start %s,token:%s\n", ctx.FullPath(), token)

	if token == "" {
		util.FailJson(ctx, cerror.ErrorTokenEmpty)
		ctx.Abort()
		return
	}

	//验证是否过期
	user, err := middleware.ParseAdminToken(token)
	if err != nil {
		util.FailJson(ctx, err)
		ctx.Abort()
		return
	}

	//如果没有过期，还需要检查是不是和系统存储的当前token一致，（如用户在其他机器上又登录了，那么之前的token就失效了）
	checkRes := service.CheckCurrentToken(user, token)
	if checkRes.Code() != cerror.ErrorUserAuthSucc.Code() {
		util.FailJson(ctx, checkRes)
		ctx.Abort()
		return
	}
	logger.Infoc(ctx, "[%s] end***  result:  check ok", "CheckAuth Controller")
	ctx.Next()

}

func AdminLogin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "Login Controller")
	// 获取参数
	var param dto.AdminLoginParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "[%s] input param: %+v", "Login Controller", param)
	}
	// 参数校验

	// 调用service
	tokenStr, res := service.Login(ctx, param.Name, param.Password)

	// 结果返回
	if res.Code() == cerror.ErrorLoginSucc.Code() {
		var loginRes dto.LoginRes
		loginRes.Token = tokenStr

		util.SuccessJson(ctx, loginRes)
	} else if res.Code() == cerror.ErrorLoginAgain.Code() {
		var loginRes dto.LoginRes
		loginRes.Token = tokenStr

		util.FailJsonData(ctx, cerror.ErrorLoginAgain, loginRes)

	} else if res.Code() == cerror.ErrorUserAuthSucc.Code() {
		var loginRes dto.LoginRes
		loginRes.Token = tokenStr

		util.SuccessJson(ctx, loginRes)

	} else {
		util.FailJson(ctx, res)
	}

	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,token =%+v", "Login Controller", res, tokenStr)
}

// /edu/v1/perm/logout
func AdminLogout(ctx *gin.Context) {

	//获取参数
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		util.FailJson(ctx, cerror.ErrorTokenEmpty)
	}
	//参数校验

	//调用service
	res := service.Logout(ctx, tokenStr)

	//结果返回

	if res == common.Success {
		util.SuccessJson(ctx, nil)
	} else {
		util.FailJson(ctx, cerror.InvalidParams)
	}
	if util.IsDebug() {
		logger.Infoc(ctx, "---->out param: %+v", res)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v", "Logout Controller", res)
}

type UserFake struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	Email            string `json:"email"`
	Job              string `json:"job"`
	JobName          string `json:"jobName"`
	Organization     string `json:"organization"`
	OrganizationName string `json:"organizationName"`
	Location         string `json:"location"`
	LocationName     string `json:"locationName"`
	Introduction     string `json:"introduction"`
	PersonalWebsite  string `json:"personalWebsite"`
	Phone            string `json:"phone"`
	RegistrationDate string `json:"registrationDate"`
	AccountId        string `json:"accountId"`
	Certification    int    `json:"certification"`

	Role int `json:"role"`
	dto.UserRes
}

func GetIdFromToken(tokenStr string) string { //

	var userIDStr string
	token, err := jwt.ParseWithClaims(tokenStr, &middleware.CustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	if claims, ok := token.Claims.(*middleware.CustomClaimsAdmin); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
		userIDStr = claims.UserName
	} else {
		fmt.Println(err)
	}

	return userIDStr

}

func GetUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start***", "GetUser Controller")
	//获取参数

	tokenStr := ctx.GetHeader("Authorization")
	loginId := GetIdFromToken(tokenStr)

	//参数校验

	var param dto.AdminParam
	param.Name = loginId
	//调用service
	user, err := service.GetAdmin(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var fake UserFake //todo：remove THIS not used
		if len(user) > 0 {
			fake.Name = user[0].Name
			fake.LoginID = user[0].Name
		}

		fake.UserID = user[0].ID
		fake.OrganizationID = 1 //todo:for test

		fake.Role = 1 // todo:role需要填上值
		fake.Avatar = "//lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png"
		fake.Email = "wangliqun@email.com"
		fake.Job = "前端艺术家"
		fake.Organization = "Frontend"
		fake.OrganizationName = "前端"
		fake.Location = "beijing"
		fake.LocationName = "北京"
		fake.Introduction = "人潇洒，性温存"
		fake.PersonalWebsite = "https://www.arco.design"
		fake.Phone = "150****0000"
		fake.RegistrationDate = "2013-05-10 12:10:00"
		fake.AccountId = "15012312300"
		fake.Certification = 1

		util.SuccessJson(ctx, fake)
	}
	logger.Infoc(ctx, "[%s] end***  result:  res=%+v,err=%+v", "GetUser Controller", user, err)
}

func Copyright(ctx *gin.Context) {
	var copyright string
	copyright = "<pre><span style=\"color: #ced4d9; font-size: 12pt;\">Copyright 2014 - 2023 IBE edu. All Rights Reserved</span></pre>"

	util.SuccessJson(ctx, copyright)
}

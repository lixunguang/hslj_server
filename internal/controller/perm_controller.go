package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lxg_jz/internal/service"
	"lxg_jz/pkg/cerror"
	"lxg_jz/pkg/util"
)

// /edu/v1/perm/login
func Login(ctx *gin.Context) {
	//参数校验
	fmt.Printf(ctx.Request.Method)
	fmt.Printf(ctx.ContentType())
	//post解析
	s, _ := ioutil.ReadAll(ctx.Request.Body) //把	body 内容读入字符串 s
	fmt.Printf("\n%s\n", string(s))

	//调用service
	res := service.Login()

	//结果返回
	if res == 0 {
		util.SuccessJson(ctx, nil)
	} else {

		util.FailJson(ctx, cerror.InvalidParams)
	}

	//logger.Infoc(ctx, "[%s] end, params:%+v", method, params)
	//c.JSON(http.StatusOK,"welcom,education edtion");
}

func Logout(ctx *gin.Context) {
	//参数校验
	fmt.Printf(ctx.Request.Method)
	fmt.Printf(ctx.ContentType())
	//post解析
	s, _ := ioutil.ReadAll(ctx.Request.Body) //把	body 内容读入字符串 s
	fmt.Printf("\n%s\n", string(s))

	res := service.Login()
	if res == 0 {
		util.SuccessJson(ctx, nil)
	} else {

		util.FailJson(ctx, cerror.InvalidParams)
	}

}

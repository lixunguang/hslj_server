package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lxg_jz/pkg/util"
)

func UserChangePassword(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserRegister(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserInfo(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserFavorite(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserHistory(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

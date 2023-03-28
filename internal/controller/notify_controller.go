package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lxg_jz/pkg/util"
)

func NotifyNewsLatest(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func NotifyNews(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

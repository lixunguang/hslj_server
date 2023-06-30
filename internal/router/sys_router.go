package router

import (
	"hslj/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sysRouter(e *gin.Engine) {
	g := e.Group("")
	g.GET("/hello", hello)
	g.POST("/sys/copyright", Copyright)
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello,华山论毽 2023")
}

func Copyright(ctx *gin.Context) {
	var copyright string
	copyright = "<pre><span style=\"color: #ced4d9; font-size: 12pt;\">Copyright 华山论毽. All Rights Reserved</span></pre>"

	util.SuccessJson(ctx, copyright)
}

package router

import (
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sysRouter(e *gin.Engine) {
	g := e.Group("")
	g.GET("/hello", hello)
	g.POST("/edu/v1/sys/copyright", Copyright)
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello,yundao edu 2023")
}

func Copyright(ctx *gin.Context) {
	var copyright string
	copyright = "<pre><span style=\"color: #ced4d9; font-size: 12pt;\">Copyright 2014 - 2023 IBE edu. All Rights Reserved</span></pre>"

	util.SuccessJson(ctx, copyright)
}

func About(ctx *gin.Context) {
	var copyright string
	copyright = "<pre><span style=\"color: #ced4d9; font-size: 12pt;\">Copyright 2014 - 2023 IBE edu. All Rights Reserved</span></pre>"

	util.SuccessJson(ctx, copyright)
}

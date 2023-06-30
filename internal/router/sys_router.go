package router

import (
<<<<<<< HEAD
	"hslj/pkg/util"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"github.com/gin-gonic/gin"
	"hslj/pkg/util"
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

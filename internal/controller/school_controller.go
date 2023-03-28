package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lxg_jz/pkg/util"
)

func SchoolDept(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func SchoolDeptClass(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func SchoolDeptClassStudent(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

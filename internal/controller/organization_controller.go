package controller

import (
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

//组织相关接口，学校，院系，班级

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

package controller

import (
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

//评分相关接口

func CourseLessonWorkScroeTable(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func CourseLessonWorkMark(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func CourseLessonWorkUp(ctx *gin.Context) {

}

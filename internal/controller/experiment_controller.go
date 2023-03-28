package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lxg_jz/pkg/util"
)

// /edu/v1/experiment/resource/recommend
func ExperimentResourceRecommend(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)
	//util.FailJson(ctx, cerror.InvalidParams)
	//ctx.JSON(102,"hello");

}

func ExperimentCategoryResource(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseOverview(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterOverview(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterContent(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterRefer(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterExperiment(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterWork(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterUpdate(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterWorkUp(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterWorkScroeTable(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func ExperimentCourseChapterWorkMark(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lxg_jz/pkg/util"
)

func LeaningCategory(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func LeaningResourceRecommend(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func LeaningCategoryResource(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func LeaningCategoryResourceRecommend(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

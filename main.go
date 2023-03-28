package main

import (
	"github.com/gin-gonic/gin"
	"lxg_jz/internal/model/db"
	router2 "lxg_jz/internal/router"
	"lxg_jz/pkg/util"
)

func main() {

	//Test()

	db.OpenDB()
	// db.GetData()

	gin.SetMode(util.GetMode())
	r := gin.Default()

	//增加访问日志、traceId
	//r.Use(middleware.AccessLogMiddleware, middleware.AddTraceId)

	//加载路由
	router2.Router(r)
	r.Run(util.GetRunPort())
}

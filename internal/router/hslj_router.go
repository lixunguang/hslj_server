package router

import (

	"github.com/gin-gonic/gin"
	"hslj/internal/controller"
)

//教育管理平台后台接口
func hsljRouter(e *gin.Engine) {
	g := e.Group("/v1")

	//权限 perm
	g.POST("/perm/login", controller.AdminLogin)
	g.POST("/copyright", controller.Copyright)
	//if !util.IsDebug() {
	//g.Use(controller.CheckAdminAuth) //middleware
	//}

	g.POST("/perm/logout", controller.AdminLogout)

	// 用户 增加 删除 查找
	g.POST("/student/add", controller.AddUser) //增加学生
	g.POST("/student/del", controller.DelUser)
	g.POST("/student/all", controller.AllUser)
	g.POST("/student/update", controller.UpdateUser)

	g.POST("/user", controller.GetUser) //后台api与平台api共用一个

	// 管理员 增加 删除 查找
	g.POST("/administrator/add", controller.AddAdmin)
	g.POST("/administrator/del", controller.DelAdmin)

	//新闻
	g.POST("/news/latest", controller.GetNewsLatest)
	g.POST("/news/all", controller.GetNewsAll)
	g.POST("/news/detail", controller.GetNews)

	//地点
	g.POST("/location", controller.GetLocation)
	g.POST("/location/list", controller.GetLocationList)
	g.POST("/location/add", controller.AddLocation)
	g.POST("/location/hot", controller.GetLocationHot)

	//地点-打卡
	g.POST("/record/add", controller.AddRecord)
	g.POST("/record/location/get", controller.GetRecordByUserID)
	g.POST("/record/user/get", controller.GetRecordByLocationID)
}

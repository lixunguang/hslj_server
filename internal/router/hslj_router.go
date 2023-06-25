package router

import (
	"edu-imp/internal/controller"
	"github.com/gin-gonic/gin"
)

//教育管理平台后台接口
func hsljRouter(e *gin.Engine) {
	g := e.Group("/admin")

	//权限 perm
	g.POST("/perm/login", controller.AdminLogin)
	g.POST("/copyright", controller.Copyright)
	//if !util.IsDebug() {
	g.Use(controller.CheckAdminAuth) //middleware
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

}
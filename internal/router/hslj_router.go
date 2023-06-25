package router

import (
	"edu-imp/internal/admin_controller"
	"github.com/gin-gonic/gin"
)

//教育管理平台后台接口
func hsljRouter(e *gin.Engine) {
	g := e.Group("/admin")

	//权限 perm
	g.POST("/perm/login", admin_controller.AdminLogin)
	g.POST("/copyright", admin_controller.Copyright)
	//if !util.IsDebug() {
	g.Use(admin_controller.CheckAdminAuth) //middleware
	//}

	g.POST("/perm/logout", admin_controller.AdminLogout)

	// 用户 增加 删除 查找
	g.POST("/student/add", admin_controller.AddUser) //增加学生
	g.POST("/student/del", admin_controller.DelUser)
	g.POST("/student/all", admin_controller.AllUser)
	g.POST("/student/update", admin_controller.UpdateUser)

	g.POST("/user", admin_controller.GetUser) //后台api与平台api共用一个

	g.POST("/teacher/add", admin_controller.AddTeacher) //增加教师
	g.POST("/teacher/del", admin_controller.DelTeacher)
	g.POST("/teacher/all", admin_controller.AllTeacher)
	g.POST("/teacher/update", admin_controller.UpdateTeacher)

	// 管理员 增加 删除 查找
	g.POST("/administrator/add", admin_controller.AddAdmin)
	g.POST("/administrator/del", admin_controller.DelAdmin)

}

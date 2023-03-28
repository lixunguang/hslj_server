package router

import (
	"github.com/gin-gonic/gin"
	"lxg_jz/internal/controller"
)

func educationRouter(e *gin.Engine) {
	g := e.Group("/edu/v1")

	//user
	// /edu/v1/user/change_password
	g.POST("/user/change_password", controller.UserChangePassword)
	// /edu/v1/user/register
	g.POST("/user/register", controller.UserRegister)
	// /edu/v1/user/info
	g.POST("/user/info", controller.UserInfo)
	// /edu/v1/user/history
	g.POST("/user/history", controller.UserHistory)
	// /edu/v1/user/favorite
	g.POST("/user/favorite", controller.UserFavorite)

	//school
	// /edu/v1/school/dept
	g.POST("/school/dept", controller.SchoolDept)
	// /edu/v1/school/dept/class
	g.POST("/school/dept/class", controller.SchoolDeptClass)
	// /edu/v1/school/dept/class/student
	g.POST("/school/dept/class/student", controller.SchoolDeptClassStudent)

	//perm
	// /edu/v1/perm/login
	g.POST("/perm/login", controller.Login)
	// /edu/v1/perm/logout
	g.POST("/perm/logout", controller.Logout)

	//notify
	// /edu/v1/notify/news/latest
	g.POST("/notify/news/latest", controller.NotifyNewsLatest)
	// /edu/v1/notify/news
	g.POST("/notify/news", controller.NotifyNews)

	//leaning
	// /edu/v1/leaning/category
	g.POST("/leaning/category", controller.LeaningCategory)
	// /edu/v1/leaning/resource/recommend
	g.POST("/leaning/resource/recommend", controller.LeaningResourceRecommend)
	// /edu/v1/leaning/category/resource
	g.POST("/leaning/category/resource", controller.LeaningCategoryResource)
	// /edu/v1/leaning/category/resource/recommend
	g.POST("/leaning/category/resource/recommend", controller.LeaningCategoryResourceRecommend)

	//experiment
	// /edu/v1/experiment/resource/recommend
	g.GET("/experiment/resource/recommend", controller.ExperimentResourceRecommend)
	// /edu/v1/experiment/category/resource
	g.POST("/experiment/category/resource", controller.ExperimentCategoryResource)
	// /edu/v1/experiment/course/overview
	g.POST("/experiment/course/overview", controller.ExperimentCourseOverview)
	// /edu/v1/experiment/course/chapter/overview
	g.POST("/experiment/course/chapter/overview", controller.ExperimentCourseChapterOverview)
	// /edu/v1/experiment/course/chapter/content
	g.POST("/experiment/course/chapter/content", controller.ExperimentCourseChapterContent)
	// /edu/v1/experiment/course/chapter/refer
	g.POST("/experiment/course/chapter/refer", controller.ExperimentCourseChapterRefer)
	// /edu/v1/experiment/course/chapter/experiment
	g.POST("/experiment/course/chapter/experiment", controller.ExperimentCourseChapterExperiment)
	// /edu/v1/experiment/course/chapter/work
	g.POST("/experiment/course/chapter/work", controller.ExperimentCourseChapterWork)
	// /edu/v1/experiment/course/chapter/update
	g.POST("/experiment/course/chapter/update", controller.ExperimentCourseChapterUpdate)
	// /edu/v1/experiment/course/chapter/work/up
	g.POST("/experiment/course/chapter/work/up", controller.ExperimentCourseChapterWorkUp)

	// /edu/v1/experiment/course/chapter/work/scroe_table
	g.POST("/experiment/course/chapter/work/scroe_table", controller.ExperimentCourseChapterWorkScroeTable)
	// /edu/v1/experiment/course/chapter/work/mark
	g.POST("/experiment/course/chapter/work/mark", controller.ExperimentCourseChapterWorkMark)

}

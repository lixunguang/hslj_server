package router

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {
	sysRouter(e)

	hsljRouter(e)
}

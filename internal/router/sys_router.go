package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sysRouter(e *gin.Engine) {
	g := e.Group("")
	g.GET("/", hello)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello,sys")
}

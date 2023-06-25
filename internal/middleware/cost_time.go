package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Function CostTime is cal time consume.
func CostTime(c *gin.Context) {

	//return func(c *gin.Context) {
	//请求前获取当前时间
	nowTime := time.Now()

	//请求处理
	c.Next()

	log.Printf("the request URL %s cost %v", c.Request.URL.String(), time.Since(nowTime))
	//}
}

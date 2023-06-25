package middleware

import (
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

//  记录响应时长
func AccessLogMiddleware(c *gin.Context) {
	beginTime := time.Now()

	c.Next()

	defer func() {
		duration := time.Now().Sub(beginTime)
		statusCode := 200
		rerr := recover()
		if rerr != nil {
			statusCode = 500
			logger.Errorf("%s panic : %+v", c.Request.URL.Path, rerr)
		} else {
			//  模拟超时记录
			if duration.Seconds() >= 20 {
				statusCode = 504
			}
		}
		accessFormat := `%s %s %s [%s] "%s %s %s" %d %s %d %d`
		logger.AccessInfof(accessFormat,
			c.ClientIP(),
			"-",
			"-",
			time.Now().Format("2006/01/02 15:04:05 -0700"),

			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			statusCode,
			"0",
			int64(duration/time.Millisecond)+1, //  加1毫秒，防止go整除舍去导致记录为0
			runtime.NumGoroutine(),
		)
	}()

}

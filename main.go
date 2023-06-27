package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"hslj/config"
	"hslj/internal/middleware"
	"hslj/internal/model/mysql"
	router2 "hslj/internal/router"
	"hslj/pkg/util"
	"io/ioutil"
	"strings"
)

func main() {

	mysql.OpenDB()
	// mysql.GetData()

	gin.SetMode(util.GetMode())
	r := gin.Default()

	//r.Use(middleware.CostTime)

	//加载静态资源，替代nginx，如果使用nginx，则去掉本部分代码

	//文件服务器1
	r.Use(static.Serve("/uploads", static.LocalFile(config.GetExecDirectory()+"/uploads", false)))
	//文件服务器2
	r.Use(static.Serve("/resource", static.LocalFile(config.GetExecDirectory()+"/resource", false)))
	//文件服务器3
	r.Use(static.Serve("/design", static.LocalFile(config.GetExecDirectory()+"/design", false)))
	//文件服务器4
	r.Use(static.Serve("/", static.LocalFile(config.GetExecDirectory()+"/dist2", false)))

	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})

	//增加访问日志、traceId

	r.Use(middleware.AccessLogMiddleware, middleware.AddTraceId)
	r.Use(middleware.Cors())
	//加载路由
	router2.Router(r)

	r.Run(util.GetRunPort())
}

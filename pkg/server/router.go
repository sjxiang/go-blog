package server

import (
	"github.com/gin-gonic/gin"
	"goblog/routes"
	"time"
	"log"
)


// 初始化 Router 路由对象
func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(Logger())
	routes.RegisterApiRoutes(r)
	
	return r
}



// 中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置变量 ava
		c.Set("ava", "向晚")

		// 请求之前 

		c.Next()

		// 请求之后
		latency := time.Since(t)
		log.Println(latency)

		// 访问我们的发送数据
		status := c.Writer.Status()
		log.Println(status)
	
	}
}


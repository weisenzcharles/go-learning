package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
路由分组
*/
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := router.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	router.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "charles")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "xiaowei")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

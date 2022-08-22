package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
创建路由
*/
func main() {
	//gin.SetMode("release")
	router := gin.Default()
	//err := route.SetTrustedProxies([]string{"192.168.0.64"})
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	router.GET("/", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s, Host: %s \n", c.ClientIP(), c.Request.Host)
	})
	router.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get!")
	})

	router.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "put!")
	})
	router.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete!")
	})
	router.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post!")
	})

	err = router.Run(":8088")
	if err != nil {
		return
	}
}

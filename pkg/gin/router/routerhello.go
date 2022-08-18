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
	//err := router.SetTrustedProxies([]string{"192.168.0.64"})
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	router.GET("/", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s, Host: %s \n", c.ClientIP(), c.Request.Host)
	})
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	err = router.Run(":8088")
	if err != nil {
		return
	}
}

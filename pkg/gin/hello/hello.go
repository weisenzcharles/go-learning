package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := SetupRouter()
	if err := router.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello charles!",
	})
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	return router
}

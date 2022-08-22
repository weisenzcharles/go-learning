package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.Use(Logger())
	route.Use(gin.Recovery())

	route.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world!")
	})

	route.Run()
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 给 Context 实例设置一个值
		context.Set("message", "test")
		// 请求前
		context.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}

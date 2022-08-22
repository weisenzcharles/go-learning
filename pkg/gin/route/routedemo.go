package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	// 获取 Query 参数
	route.GET("/user", func(context *gin.Context) {
		name := context.Query("name")
		context.String(http.StatusOK, "hello %s", name)
	})
	// 获取 Post 参数
	route.POST("/user/add", func(context *gin.Context) {
		name := context.PostForm("name")
		context.String(http.StatusOK, "add user %s", name)
	})

	// GET 和 POST 混合
	route.POST("/posts", func(context *gin.Context) {
		id := context.Query("id")
		pageIndex := context.DefaultQuery("pageIndex", "0") // 设置默认值
		username := context.PostForm("username")
		password := context.DefaultPostForm("username", "") // 设置默认值

		context.JSON(http.StatusOK, gin.H{
			"id":        id,
			"pageIndex": pageIndex,
			"username":  username,
			"password":  password,
		})
	})

	// 字典参数
	route.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向
	route.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})

	// 上传文件
	route.POST("/upload", func(context *gin.Context) {
		file, err := context.MultipartForm()
		context.String(http.StatusOK, "%s uploaded %s!", file.File, err)
	})

	// Html 模版
	type user struct {
		Name string
		Age  int8
	}
	fmt.Println(filepath.Join(os.Getenv("GOPATH")))
	route.LoadHTMLGlob("templates/*")
	// route.LoadHTMLGlob("templates/*")

	user1 := &user{Name: "Charles", Age: 20}
	user2 := &user{Name: "Xiaomi", Age: 22}
	route.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.tmpl", gin.H{
			"title": "user",
			"users": [2]*user{user1, user2},
		})
	})
	route.Run()
}

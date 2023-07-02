package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 默认路由
	rt := gin.Default()

	// 指定get 请求 json路径，调用指定函数
	rt.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	rt.GET("/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/b"
		rt.HandleContext(ctx)
	})

	rt.GET("/b", func(ctx *gin.Context) {
		fmt.Println("hello i am in b")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": ctx.Request.URL,
		})
	})

	// 默认0.0.0.1：8080 run
	rt.Run()

}

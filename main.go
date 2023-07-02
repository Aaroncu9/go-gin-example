package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Info string `json:"info"`
}

func main() {
	// 默认路由
	rt := gin.Default()

	// 指定get 请求 json路径，调用指定函数
	rt.GET("/student", func(ctx *gin.Context) {
		name := ctx.Query("name")
		age := ctx.Query("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	// 默认0.0.0.1：8080 run
	rt.Run()

}

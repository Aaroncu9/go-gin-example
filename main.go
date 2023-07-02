package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func main() {
	// 默认路由
	rt := gin.Default()

	// 指定get 请求 json路径，调用指定函数
	rt.GET("/user", func(ctx *gin.Context) {
		// get request
		var u user
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			fmt.Printf("user info %#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}

	})

	// 默认0.0.0.1：8080 run
	rt.Run()

}

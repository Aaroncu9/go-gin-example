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
	rt := gin.Default()
	rt.GET("/name", func(ctx *gin.Context) {
		stu := student{
			Name: "liuming",
			Age:  20,
			Info: "三好学生",
		}
		ctx.JSON(http.StatusOK, stu)
	})

	rt.Run()

}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello gin",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user", helloHandler)
	return r
}

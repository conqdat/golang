package main

import (
	"golang_crud_gin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	routes.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

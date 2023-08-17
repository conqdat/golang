package main

import (
	"GoLang/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Name": "Cong Dat",
		})
	})
	r.GET("/people", func(context *gin.Context) {
		context.JSON(http.StatusOK, repo.ListPeople())
	})
	err := r.Run()
	if err != nil {
		return
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var r = gin.Default() // set gin routes

	r.Use(testGlobalMiddleWare) // use middleware

	r.MaxMultipartMemory = 8 << 20

	r.GET("/", homeHandler)
	r.POST("/", postHandler)
	r.GET("api/detail/:id", detailHandler)

	// Using Group in API
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", homeHandler)
			v1.POST("/", postHandler)
			v1.GET("detail/:id", detailHandler)
			v1.POST("/upload", handlerUploadImage)
		}
	}

	err := r.Run(":333") // listen port 333
	if err != nil {
		return
	}
}

func handlerUploadImage(context *gin.Context) {
	file, _ := context.FormFile("file")
	log.Println("File Name ", file.Filename)

	err := context.SaveUploadedFile(file, "images")
	context.String(http.StatusOK, "Uploaded")

	if err != nil {
		return
	}
}

func postHandler(context *gin.Context) {
	data := context.DefaultPostForm("name", "no name")
	context.JSON(http.StatusCreated, data)
}

func homeHandler(context *gin.Context) {
	//context.String(http.StatusOK, "Hello") // return String
	//context.JSON(http.StatusOK, "Json return") // return Json
	name := context.DefaultQuery("name", "Guest") // get Query Params
	age := context.DefaultQuery("age", "18")
	context.JSON(http.StatusOK, age+name)
}

func detailHandler(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, id)
}

func testGlobalMiddleWare(context *gin.Context) {
	log.Println("in mid")
	context.Next()
}

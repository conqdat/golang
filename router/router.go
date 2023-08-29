package router

import (
	"golang_crud_gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello Dat")
	})

	baseRouter := router.Group("/api")

	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Udpate)
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}

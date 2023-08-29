package controller

import (
	"golang_crud_gin/helper"
	"golang_crud_gin/request"
	"golang_crud_gin/response"
	"golang_crud_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TagController struct {
	tagService service.TagService
}

func NewTagController(service service.TagService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	log.Info().Msg("Find ALL ")

	tags := controller.tagService.FindAll()

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tags,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, apiResponse)
}

func (controller *TagController) Create(ctx *gin.Context) {
	log.Info().Msg("CREATE TAG")
	createTagRequest := request.CreateTagREquest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)
	apiResponse := response.ApiResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, apiResponse)
}

func (controller *TagController) Udpate(ctx *gin.Context) {
	log.Info().Msg("UPDATE TAG")
	updateTagRequest := request.UpdateTagRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id

	controller.tagService.Update(updateTagRequest)

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, apiResponse)
}

func (controller *TagController) FindById(ctx *gin.Context) {
	log.Info().Msg("Find Tag")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tag := controller.tagService.FindById(id)

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tag,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, apiResponse)
}

func (controller *TagController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete Tag")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagService.Delete(id)

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, apiResponse)
}

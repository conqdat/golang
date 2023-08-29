package main

import (
	"golang_crud_gin/config"
	"golang_crud_gin/controller"
	"golang_crud_gin/helper"
	"golang_crud_gin/model"
	"golang_crud_gin/repository"
	"golang_crud_gin/router"
	"golang_crud_gin/service"
	"net/http"

	"github.com/go-delve/delve/pkg/config"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Server is running at: 8888")

	// DB
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repo
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

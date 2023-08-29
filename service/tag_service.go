package service

import (
	"golang_crud_gin/request"
	"golang_crud_gin/response"
)

type TagService interface {
	Create(tags request.CreateTagRequest)
	Update(tags request.UpdateTagRequest)
	Delete(tagId int)
	FindById(tagsId int) response.TagResponse
	FindAll() []response.TagResponse
}

package service

import (
	"golang_crud_gin/helper"
	"golang_crud_gin/model"
	"golang_crud_gin/repository"
	"golang_crud_gin/request"
	"golang_crud_gin/response"

	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

// Create implements TagService.
func (t *TagServiceImpl) Create(tags request.CreateTagRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagService.
func (t *TagServiceImpl) Delete(tagId int) {
	t.TagsRepository.Delete(tagId)
}

// FindAll implements TagService.
func (t *TagServiceImpl) FindAll() []response.TagResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagResponse
	for _, value := range result {
		tag := response.TagResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagService.
func (t *TagServiceImpl) FindById(tagsId int) response.TagResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagReponse := response.TagResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagReponse
}

// Update implements TagService.
func (t *TagServiceImpl) Update(tags request.UpdateTagRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)

	tags.Name = tagData.Name
	t.TagsRepository.Update(tagData)
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

package services

import (
	"volt/data/entitys"
	"volt/data/model"
	"volt/data/repository"
)

type TagService struct {
	tagRepository *repository.TagRepository
}

func NewTagService(tagRepository *repository.TagRepository) *TagService {
	return &TagService{tagRepository:tagRepository}
}


func (this *TagService) GetByID(id int64) (*entitys.Tag, error) {
	return this.tagRepository.GetById(id)
}

func (this *TagService)GetListByPage(search model.TagSearch)(tagList []entitys.Tag,total int64,err error) {
	return this.tagRepository.SearchByPage(search)
}

func (this *TagService)AddTag(name string) bool {
	return this.tagRepository.Add(name)
}

func (this *TagService)UpdateTag(tagId int64, tagName string) bool {
	return this.tagRepository.Update(tagId,tagName)
}
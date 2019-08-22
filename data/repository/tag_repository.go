package repository

import (
	"github.com/go-xorm/xorm"
	"volt/data/entitys"
	"volt/data/model"
)

type TagRepository struct {
	db *xorm.Engine
}

func NewTagRepository(db *xorm.Engine) *TagRepository {
	return &TagRepository{db}
}

func (this *TagRepository) GetById(Id int64) (*entitys.Tag,error) {
	tag := new(entitys.Tag)
	_, err := this.db.Id(Id).Get(tag)
	return tag,err
}


func (this *TagRepository)Add(name string) bool  {
	tag := new(entitys.Tag)
	tag.Name = name
	_, err := this.db.Insert(tag)
	if err != nil {
		return false
	}
	return true
}

func (this *TagRepository)Update(tagId int64,tagName string) bool  {
	tag := new(entitys.Tag)
	tag.Name = tagName
	_, err := this.db.Id(tagId).Update(tag)
	if err != nil {
		return false
	}
	return true
}

func (this *TagRepository) SearchByPage(search model.TagSearch) ([]entitys.Tag, int64, error) {
	tag := new(entitys.Tag)
	total, err := this.db.Count(tag)
	var tags []entitys.Tag

	var order string
	if search.Isdesc {
		order = search.OrderBy + " desc"
	}else {
		order =search.OrderBy +" asc"
	}

	err = this.db.OrderBy(order).Limit(search.PageSize,search.PageIndex *search.PageSize).Find(&tags)

	return tags, total, err
}

func (this *TagRepository) GetAll()(tags []entitys.Tag) {
	this.db.OrderBy("CreateTime desc").Find(&tags)
    return tags
}
package repository

import (
	"github.com/go-xorm/xorm"
	"time"
	"volt/data/entitys"
)

type TagArticleRepository struct {
	db *xorm.Engine
}

func NewTagArticleRepository(db *xorm.Engine) *TagArticleRepository {
	return &TagArticleRepository{db}
}

func (this *TagArticleRepository) GetById(Id int64) (*entitys.TagArticle,error) {
	tagpost := new(entitys.TagArticle)
	_, err := this.db.Id(Id).Get(tagpost)
	return tagpost,err
}


func (this *TagArticleRepository)Add(tagid int,articleid int64) bool  {
	articletag := new(entitys.TagArticle)
	articletag.TagId = tagid
	articletag.ArticleId = articleid
	articletag.ArticleTime = time.Now()
	_, err := this.db.Insert(articletag)
	if err != nil {
		return false
	}
	return true
}

func (this *TagArticleRepository)Update(tagid int,articleid int64) bool  {
	articletag := new(entitys.TagArticle)
	articletag.TagId = tagid
	articletag.ArticleId = articleid
	_, err := this.db.Id(articleid).Update(articletag)
	if err != nil {
		return false
	}
	return true
}

package repository

import (
	"github.com/go-xorm/xorm"
	"volt/data/entitys"
	"volt/data/model"
)

type ArticleRepository struct {
	db *xorm.Engine
}

func NewArticleRepository(db *xorm.Engine) *ArticleRepository {
	return &ArticleRepository{db}
}

//根据Id获取文章信息
func (this *ArticleRepository) GetById(Id int32) (*entitys.Article,error) {
	post := new(entitys.Article)
	_, err := this.db.Id(Id).Get(post)
	return post,err
}

func (this *ArticleRepository) AddArticle(article *entitys.Article) int64 {
	id, err := this.db.Insert(article)
	if err != nil {
		return 0
	}
	return id
}

func (this *ArticleRepository) SearchByPage(search model.ArticleSearch) ([]entitys.Article, int64, error) {
	article := new(entitys.Article)
	total, err := this.db.Count(article)
	var articles []entitys.Article

	var order string
	if search.Isdesc {
		order = search.OrderBy + " desc"
	}else {
		order =search.OrderBy +" asc"
	}

	err = this.db.OrderBy(order).Limit(search.PageSize,search.PageIndex *search.PageSize).Find(&articles)

	return articles, total, err
}


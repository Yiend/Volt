package services

import (
	"volt/data/entitys"
	"volt/data/model"
	"volt/data/repository"
)

type ArticleService struct {
	articleRepository *repository.ArticleRepository
	tagRepository *repository.TagRepository
	tagArticleRepository *repository.TagArticleRepository
}

func NewArticleService(articleRepository *repository.ArticleRepository,agRepository *repository.TagRepository,tagArticleRepository *repository.TagArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepository:articleRepository,
		tagRepository:agRepository,
		tagArticleRepository:tagArticleRepository,
	}
}

func (this *ArticleService) AddArticle(tagIds []int ,article *entitys.Article) bool {
	articleId := this.articleRepository.AddArticle(article)
	if articleId > 0 {
		for tagId := range tagIds {
            this.tagArticleRepository.Add(tagId,articleId)
		}
		return true
	}
	return  false
}

func (this *ArticleService) GetByID(id int32) (*entitys.Article, error) {
	return this.articleRepository.GetById(id)
}

func (this *ArticleService)GetListByPage(search model.ArticleSearch)(articleList []entitys.Article,total int64,err error) {
   return this.articleRepository.SearchByPage(search)
}

func (this *ArticleService) GetTags()map[int32]string  {
	tags := make(map[int32]string)
	tagList := this.tagRepository.GetAll()
	for _,v := range tagList {
		tags[v.ID] =v.Name
	}
	return tags
}
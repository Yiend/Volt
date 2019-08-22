package viewmodel

type ArticleViewModel struct {
	Title         string                `json:"title" validate:"required"`
	Author        string                `json:"author" validate:"required"`
	TitleColor    string                `json:"titlecolor"`
	Tags          []TagViewModel        `json:"tags" validate:"required"`
	Content       string                `json:"content" validate:"required"`
}

type TagViewModel struct {
	ID     int      `json:"id" validate:"required"`
	Name   string    `json:"name" validate:"required"`
}
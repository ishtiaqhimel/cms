package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
)

func articleCreateReqToArticle(req *model.ArticleCreateReq) *entity.Article {
	return &entity.Article{
		Title:      req.Title,
		Body:       req.Body,
		Status:     entity.ArticleStatusDraft,
		CategoryID: req.CategoryID,
		AuthorID:   req.AuthorID,
	}
}

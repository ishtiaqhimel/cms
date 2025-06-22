package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/repository"
)

type articleUsecase struct {
	repo repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		repo: repo,
	}
}

type ArticleUsecase interface {
	CreateArticle(req *model.ArticleCreateReq) (*model.EntityID, error)
}

func (a *articleUsecase) CreateArticle(req *model.ArticleCreateReq) (*model.EntityID, error) {
	article := articleCreateReqToArticle(req)
	// TODO: check if category_id and author_id is valid
	return a.repo.CreateArticle(article)
}

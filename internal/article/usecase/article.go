package usecase

import (
	"context"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/repository"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

type articleUsecase struct {
	repo   repository.ArticleRepository
	userUC usecase.UserUsecase
}

func NewArticleUsecase(repo repository.ArticleRepository, userUC usecase.UserUsecase) ArticleUsecase {
	return &articleUsecase{
		repo:   repo,
		userUC: userUC,
	}
}

type ArticleUsecase interface {
	CreateArticle(ctx context.Context, req *model.ArticleCreateReq) (*model.EntityID, error)
}

func (a *articleUsecase) CreateArticle(ctx context.Context, req *model.ArticleCreateReq) (*model.EntityID, error) {
	article := articleCreateReqToArticle(req)

	author, err := a.userUC.GetUserByID(article.AuthorID)
	if err != nil {
		return nil, err
	}

	if !utils.DefaultRBAC(ctx, author.Role).
		IsAuthorized() {
		return nil, response.ErrUnauthorized
	}

	return a.repo.CreateArticle(article)
}

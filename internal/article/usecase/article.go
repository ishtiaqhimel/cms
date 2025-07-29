package usecase

import (
	"context"
	"errors"

	"gorm.io/gorm"

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
	GetArticleByID(id string) (*model.ArticleResp, error)
	UpdateArticle(id string, req *model.ArticleUpdateReq) error
	ListArticleByFilter(filter *ArticleListFilter, pg *utils.Pagination) ([]*model.ArticleResp, int64, error)
	DeleteArticleByID(id string) error
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

func (a *articleUsecase) GetArticleByID(id string) (*model.ArticleResp, error) {
	article, err := a.repo.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return toArticleResp(article), nil
}

func (a *articleUsecase) UpdateArticle(id string, req *model.ArticleUpdateReq) error {
	article, updateCols := articleUpdateReqToArticle(id, req)
	return a.repo.UpdateArticle(article, updateCols)
}

func (a *articleUsecase) ListArticleByFilter(filter *ArticleListFilter, pg *utils.Pagination) ([]*model.ArticleResp, int64, error) {
	f := &repository.ArticleListFilter{
		ID:         filter.ID,
		Status:     filter.Status,
		CategoryID: filter.CategoryID,
		AuthorID:   filter.AuthorID,
	}

	res, total, err := a.repo.ListArticleByFilter(f, pg)
	if err != nil {
		return nil, 0, err
	}

	return toArticleRespList(res), total, nil
}

func (a *articleUsecase) DeleteArticleByID(id string) error {
	return a.repo.DeleteArticleByID(id)
}

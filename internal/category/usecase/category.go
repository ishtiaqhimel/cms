package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/repository"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

type CategoryUsecase interface {
	CreateCategory(req *model.CategoryCreateReq) (*model.EntityID, error)
}

func (a *categoryUsecase) CreateCategory(req *model.CategoryCreateReq) (*model.EntityID, error) {
	category := categoryCreateReqToCategory(req)

	slug, err := utils.GenerateSlug(category.Name)
	if err != nil {
		return nil, err
	}
	category.Slug = slug

	return a.repo.CreateCategory(category)
}

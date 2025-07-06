package usecase

import (
	"errors"

	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/repository"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
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
	GetCategoryByID(id string) (*model.CategoryResp, error)
	UpdateCategory(id string, req *model.CategoryUpdateReq) error
	ListCategoryByFilter(filter *CategoryListFilter, pg *utils.Pagination) ([]*model.CategoryResp, int64, error)
	DeleteCategoryByID(id string) error
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

func (a *categoryUsecase) GetCategoryByID(id string) (*model.CategoryResp, error) {
	category, err := a.repo.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return toCategoryResp(category), nil
}

func (a *categoryUsecase) UpdateCategory(id string, req *model.CategoryUpdateReq) error {
	category, updateCols := categoryUpdateReqToCategory(id, req)
	return a.repo.UpdateCategory(category, updateCols)
}

func (a *categoryUsecase) ListCategoryByFilter(filter *CategoryListFilter, pg *utils.Pagination) ([]*model.CategoryResp, int64, error) {
	f := &repository.CategoryListFilter{
		ID:       filter.ID,
		Name:     filter.Name,
		ParentID: filter.ParentID,
		IsActive: filter.IsActive,
	}

	res, total, err := a.repo.ListCategoryByFilter(f, pg)
	if err != nil {
		return nil, 0, err
	}

	return toCategoryRespList(res), total, nil
}

func (a *categoryUsecase) DeleteCategoryByID(id string) error {
	return a.repo.DeleteCategoryByID(id)
}

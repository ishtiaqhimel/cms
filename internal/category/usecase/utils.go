package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
)

type CategoryListFilter struct {
	ID       *string
	Name     *string
	ParentID *string
	IsActive *bool
}

func categoryCreateReqToCategory(req *model.CategoryCreateReq) *entity.Category {
	return &entity.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		IsActive:    req.IsActive,
	}
}

func categoryUpdateReqToCategory(id string, req *model.CategoryUpdateReq) (*entity.Category, []string) {
	category := &entity.Category{
		ID: id,
	}
	var updateCols []string

	if req.Name != nil {
		category.Name = *req.Name
		updateCols = append(updateCols, "name")
	}

	if req.Description != nil {
		category.Description = req.Description
		updateCols = append(updateCols, "description")
	}

	if req.ParentID != nil {
		category.ParentID = req.ParentID
		updateCols = append(updateCols, "parent_id")
	}

	if req.IsActive != nil {
		category.IsActive = *req.IsActive
		updateCols = append(updateCols, "is_active")
	}

	return category, updateCols
}

func toCategoryResp(category *entity.Category) *model.CategoryResp {
	if category == nil {
		return nil
	}

	return &model.CategoryResp{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		ParentID:    category.ParentID,
		IsActive:    category.IsActive,
	}
}

func toCategoryRespList(in []*entity.Category) []*model.CategoryResp {
	out := make([]*model.CategoryResp, 0)

	for _, item := range in {
		temp := toCategoryResp(item)
		if temp != nil {
			out = append(out, temp)
		}
	}

	return out
}

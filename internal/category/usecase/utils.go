package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
)

func categoryCreateReqToCategory(req *model.CategoryCreateReq) *entity.Category {
	return &entity.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		IsActive:    req.IsActive,
	}
}

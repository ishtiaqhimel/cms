package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
)

type categoryRepository struct {
	*gorm.DB
}

// NewCategoryRepository will create an object that represent the CategoryRepository interface
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*model.EntityID, error)
}

func (a *categoryRepository) CreateCategory(category *entity.Category) (*model.EntityID, error) {
	category.ID = uuid.New().String()
	resp := a.DB.Table(category.TableName()).Create(category)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &model.EntityID{
		ID: category.ID,
	}, nil
}

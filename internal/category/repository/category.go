package repository

import (
	"time"

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
	currTime := time.Now().UTC()
	category.CreatedAt = currTime
	category.UpdatedAt = currTime

	err := a.DB.Table(category.TableName()).Create(category).Error
	if err != nil {
		return nil, err
	}

	return &model.EntityID{
		ID: category.ID,
	}, nil
}

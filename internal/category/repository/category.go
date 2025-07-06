package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
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
	GetCategoryByID(id string) (*entity.Category, error)
	UpdateCategory(category *entity.Category, updatedCols []string) error
	ListCategoryByFilter(filter *CategoryListFilter, pg *utils.Pagination) ([]*entity.Category, int64, error)
	DeleteCategoryByID(id string) error
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

func (a *categoryRepository) GetCategoryByID(id string) (*entity.Category, error) {
	category := &entity.Category{ID: id}
	resp := a.DB.Table(category.TableName()).First(category)

	return category, resp.Error
}

func (a *categoryRepository) UpdateCategory(category *entity.Category, updatedCols []string) error {
	resp := a.DB.Table(category.TableName()).Select(updatedCols).Updates(category)

	return resp.Error
}

func (a *categoryRepository) ListCategoryByFilter(filter *CategoryListFilter, pg *utils.Pagination) ([]*entity.Category, int64, error) {
	categories := make([]*entity.Category, 0)

	tx := a.DB.Table(entity.Category{}.TableName())

	if filter.ID != nil {
		tx = tx.Where("id = ?", *filter.ID)
	}

	if filter.Name != nil {
		tx = tx.Where("name ~* ?", *filter.Name)
	}

	if filter.ParentID != nil {
		tx = tx.Where("parent_id = ?", *filter.ParentID)
	}

	if filter.IsActive != nil {
		tx = tx.Where("is_active = ?", *filter.IsActive)
	}

	totalRecords := int64(0)
	if resp := tx.Where("deleted_at IS NULL").Count(&totalRecords); resp.Error != nil {
		return nil, 0, resp.Error
	}

	offset := (pg.Page - 1) * pg.PageSize
	if resp := tx.Limit(pg.PageSize).Offset(offset).Find(&categories); resp.Error != nil {
		return nil, 0, resp.Error
	}

	return categories, totalRecords, nil
}

func (a *categoryRepository) DeleteCategoryByID(id string) error {
	resp := a.DB.Table(entity.Category{}.TableName()).Delete(&entity.Category{ID: id})

	return resp.Error
}

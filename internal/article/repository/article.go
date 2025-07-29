package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

type articleRepository struct {
	*gorm.DB
}

// NewArticleRepository will create an object that represent the ArticleRepository interface
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

type ArticleRepository interface {
	CreateArticle(article *entity.Article) (*model.EntityID, error)
	GetArticleByID(id string) (*entity.Article, error)
	UpdateArticle(article *entity.Article, updateCols []string) error
	ListArticleByFilter(filter *ArticleListFilter, pg *utils.Pagination) ([]*entity.Article, int64, error)
	DeleteArticleByID(id string) error
}

func (a *articleRepository) CreateArticle(article *entity.Article) (*model.EntityID, error) {
	article.ID = uuid.New().String()
	resp := a.DB.Table(article.TableName()).Create(article)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &model.EntityID{
		ID: article.ID,
	}, nil
}

func (a *articleRepository) GetArticleByID(id string) (*entity.Article, error) {
	article := &entity.Article{ID: id}
	resp := a.DB.Table(article.TableName()).First(article)

	return article, resp.Error
}

func (a *articleRepository) UpdateArticle(article *entity.Article, updateCols []string) error {
	resp := a.DB.Table(article.TableName()).Select(updateCols).Updates(article)

	return resp.Error
}

func (a *articleRepository) ListArticleByFilter(filter *ArticleListFilter, pg *utils.Pagination) ([]*entity.Article, int64, error) {
	articles := make([]*entity.Article, 0)

	tx := a.DB.Table(entity.Article{}.TableName()).Order("created_at desc")

	if filter.ID != nil {
		tx = tx.Where("id = ?", *filter.ID)
	}

	if filter.Status != nil {
		tx = tx.Where("status = ?", *filter.Status)
	}

	if filter.CategoryID != nil {
		tx = tx.Where("category_id = ?", *filter.CategoryID)
	}

	if filter.AuthorID != nil {
		tx = tx.Where("author_id = ?", *filter.AuthorID)
	}

	totalRecords := int64(0)
	if resp := tx.Where("deleted_at IS NULL").Count(&totalRecords); resp.Error != nil {
		return nil, 0, resp.Error
	}

	offset := (pg.PageSize - 1) * pg.PageSize
	if resp := tx.Limit(pg.PageSize).Offset(offset).Find(&articles); resp.Error != nil {
		return nil, 0, resp.Error
	}

	return articles, totalRecords, nil
}

func (a *articleRepository) DeleteArticleByID(id string) error {
	resp := a.DB.Table(entity.Article{}.TableName()).Delete(&entity.Article{ID: id})

	return resp.Error
}

package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
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

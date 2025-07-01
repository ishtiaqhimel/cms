package entity

import (
	"time"

	"database/sql/driver"
	"gorm.io/gorm"
)

type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusPublished ArticleStatus = "published"
)

func (r *ArticleStatus) Scan(value interface{}) error {
	*r = ArticleStatus(value.(string))
	return nil
}

func (r ArticleStatus) Value() (driver.Value, error) {
	return string(r), nil
}

type Article struct {
	ID         string        `gorm:"type:uuid;primaryKey"`
	Title      string        `gorm:"type:varchar(256);not null"`
	Body       string        `gorm:"not null"`
	Status     ArticleStatus `gorm:"type:article_status;default:draft"`
	CategoryID string        `gorm:"type:uuid"`
	AuthorID   string        `gorm:"type:uuid"`
	PublishAt  *time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt
}

func (a Article) TableName() string {
	return "articles"
}

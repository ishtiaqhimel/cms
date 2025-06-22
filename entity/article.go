package entity

import (
	"time"

	"database/sql/driver"
)

type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusPublished ArticleStatus = "published"
)

func (r *ArticleStatus) Scan(value interface{}) error {
	*r = ArticleStatus(value.([]byte))
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
	DeletedAt  *time.Time
}

func (a Article) TableName() string {
	return "articles"
}

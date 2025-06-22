package entity

import "time"

type Category struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null;unique"`
	Slug        string    `gorm:"type:varchar(100);not null;unique"`
	Description *string   `gorm:"type:text"`
	ParentID    *string   `gorm:"type:uuid"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time
}

func (a Category) TableName() string {
	return "categories"
}

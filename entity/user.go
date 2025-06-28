package entity

import (
	"database/sql/driver"
	"time"
)

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleEditor   UserRole = "editor"
	UserRoleReporter UserRole = "reporter"
	UserRoleReader   UserRole = "reader"
)

func (r *UserRole) Scan(src interface{}) error {
	*r = UserRole(src.(string))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(32);not null"`
	Email     string    `gorm:"type:varchar(320);unique,not null"`
	Role      UserRole  `gorm:"type:user_role;default:'reader'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time
}

func (a User) TableName() string {
	return "users"
}

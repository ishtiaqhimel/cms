package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
)

type userRepository struct {
	*gorm.DB
}

// NewUserRepository will create an object that represent the UserRepository interface
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

type UserRepository interface {
	CreateUser(user *entity.User) (*model.EntityID, error)
}

func (a *userRepository) CreateUser(user *entity.User) (*model.EntityID, error) {
	user.ID = uuid.New().String()
	currTime := time.Now().UTC()
	user.CreatedAt = currTime
	user.UpdatedAt = currTime

	err := a.DB.Table(user.TableName()).Create(user).Error
	if err != nil {
		return nil, err
	}

	return &model.EntityID{
		ID: user.ID,
	}, nil
}

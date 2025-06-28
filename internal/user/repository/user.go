package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
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
	GetUserByID(userID string) (*entity.User, error)
	UpdateUser(user *entity.User, updateCols []string) error
	ListUserByFilter(filter *UserListFilter, pg *utils.Pagination) ([]*entity.User, int64, error)
	DeleteUserByID(userID string) error
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

func (a *userRepository) GetUserByID(id string) (*entity.User, error) {
	user := &entity.User{}
	resp := a.DB.Table(user.TableName()).Where("id = ? AND deleted_at IS NULL", id).Find(user)
	return user, resp.Error
}

func (a *userRepository) UpdateUser(user *entity.User, updatedCols []string) error {
	user.UpdatedAt = time.Now().UTC()

	resp := a.DB.Table(user.TableName()).Where("deleted_at IS NULL").Select(updatedCols).Updates(user)

	return resp.Error
}

func (a *userRepository) ListUserByFilter(filter *UserListFilter, pg *utils.Pagination) ([]*entity.User, int64, error) {
	users := make([]*entity.User, 0)

	tx := a.DB.Table(entity.User{}.TableName()).Where("deleted_at IS NULL")

	if filter.ID != nil {
		tx = tx.Where("id = ?", *filter.ID)
	}

	if filter.Name != nil {
		tx = tx.Where("name ~* ?", *filter.Name)
	}

	if filter.Email != nil {
		tx = tx.Where("email = ?", *filter.Email)
	}

	if filter.Role != nil {
		tx = tx.Where("role = ?", *filter.Role)
	}

	totalRecords := int64(0)
	if resp := tx.Count(&totalRecords); resp.Error != nil {
		return nil, 0, resp.Error
	}

	offset := (pg.Page - 1) * pg.PageSize
	if resp := tx.Limit(pg.PageSize).Offset(offset).Find(&users); resp.Error != nil {
		return nil, 0, resp.Error
	}

	return users, totalRecords, nil
}

func (a *userRepository) DeleteUserByID(id string) error {
	currTime := time.Now().UTC()

	user := &entity.User{
		ID:        id,
		DeletedAt: &currTime,
	}
	resp := a.DB.Table(user.TableName()).Select("deleted_at").Updates(user)

	return resp.Error
}

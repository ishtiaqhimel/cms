package usecase

import (
	"errors"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"

	"gorm.io/gorm"

	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/repository"
)

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

type UserUsecase interface {
	CreateUser(req *model.UserCreateReq) (*model.EntityID, error)
	GetUserByID(id string) (*model.UserResp, error)
	UpdateUser(id string, req *model.UserUpdateReq) error
	ListUserByFilter(filter *UserListFilter, pg *utils.Pagination) ([]*model.UserResp, int64, error)
	DeleteUserByID(id string) error
}

func (a *userUsecase) CreateUser(req *model.UserCreateReq) (*model.EntityID, error) {
	user := userCreateReqToUser(req)
	return a.repo.CreateUser(user)
}

func (a *userUsecase) GetUserByID(id string) (*model.UserResp, error) {
	user, err := a.repo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return toUserResp(user), nil
}

func (a *userUsecase) UpdateUser(id string, req *model.UserUpdateReq) error {
	user, updateCols := userUpdateReqToUser(id, req)
	return a.repo.UpdateUser(user, updateCols)
}

func (a *userUsecase) ListUserByFilter(filter *UserListFilter, pg *utils.Pagination) ([]*model.UserResp, int64, error) {
	f := &repository.UserListFilter{
		ID:    filter.ID,
		Name:  filter.Name,
		Email: filter.Email,
		Role:  filter.Role,
	}

	res, total, err := a.repo.ListUserByFilter(f, pg)
	if err != nil {
		return nil, 0, err
	}

	return toUserRespList(res), total, nil
}

func (a *userUsecase) DeleteUserByID(id string) error {
	return a.repo.DeleteUserByID(id)
}

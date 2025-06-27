package usecase

import (
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
}

func (a *userUsecase) CreateUser(req *model.UserCreateReq) (*model.EntityID, error) {
	user := userCreateReqToUser(req)
	return a.repo.CreateUser(user)
}

package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/internal/system/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/system/repository"
)

type SystemUsecase interface {
	GetHealth() (*model.SystemHealthResp, error)
	GetTime() *model.SystemTimeResp
}

type systemUsecase struct {
	repo repository.SystemRepository
}

func NewSystemUsecase(repo repository.SystemRepository) SystemUsecase {
	return &systemUsecase{
		repo: repo,
	}
}

func (u *systemUsecase) GetHealth() (*model.SystemHealthResp, error) {
	dbOnline, err := u.repo.DBCheck()
	resp := model.SystemHealthResp{
		DBOnline: dbOnline,
	}
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (u *systemUsecase) GetTime() *model.SystemTimeResp {
	currentTime := u.repo.CurrentTime()
	resp := model.SystemTimeResp{
		CurrentTimeUnix: currentTime,
	}

	return &resp
}

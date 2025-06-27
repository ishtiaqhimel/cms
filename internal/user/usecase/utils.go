package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
)

func userCreateReqToUser(req *model.UserCreateReq) *entity.User {
	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  entity.UserRoleReader,
	}

	if req.Role != nil {
		user.Role = entity.UserRole(*req.Role)
	}

	return user
}

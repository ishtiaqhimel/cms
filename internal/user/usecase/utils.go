package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
)

type UserListFilter struct {
	ID    *string
	Name  *string
	Email *string
	Role  *string
}

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

func userUpdateReqToUser(id string, req *model.UserUpdateReq) (*entity.User, []string) {
	user := &entity.User{
		ID: id,
	}
	var updateCols []string

	if req.Name != nil {
		user.Name = *req.Name
		updateCols = append(updateCols, "name")
	}

	if req.Email != nil {
		user.Email = *req.Email
		updateCols = append(updateCols, "email")
	}

	if req.Role != nil {
		user.Role = entity.UserRole(*req.Role)
		updateCols = append(updateCols, "role")
	}

	return user, updateCols
}

func toUserResp(user *entity.User) *model.UserResp {
	if user == nil {
		return nil
	}

	return &model.UserResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}
}

func toUserRespList(in []*entity.User) []*model.UserResp {
	out := make([]*model.UserResp, 0)

	for _, item := range in {
		temp := toUserResp(item)
		if temp != nil {
			out = append(out, temp)
		}
	}

	return out
}

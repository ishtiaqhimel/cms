package model

import (
	"fmt"
	"strings"

	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

// Request Models

type UserCreateReq struct {
	Name  string  `json:"name" validate:"required"`
	Email string  `json:"email" validate:"required"`
	Role  *string `json:"role"`
}

func (req *UserCreateReq) Validate() error {
	return validateUserRole(req.Role)
}

type UserUpdateReq struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Role  *string `json:"role"`
}

func (req *UserUpdateReq) Validate() error {
	return validateUserRole(req.Role)
}

func validateUserRole(role *string) error {
	if role == nil {
		return nil
	}

	userRole := entity.UserRole(strings.ToLower(*role))
	if userRole == entity.UserRoleAdmin {
		return fmt.Errorf("user role '%s' cannot be created/updated", *role)
	}

	if !utils.ValueInSlice(userRole, []entity.UserRole{entity.UserRoleEditor, entity.UserRoleReporter, entity.UserRoleReader}) {
		return fmt.Errorf("invalid user role '%s'", *role)
	}

	return nil
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

type UserResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

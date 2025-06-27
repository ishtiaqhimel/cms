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
	if req.Role == nil {
		return nil
	}

	userRole := entity.UserRole(strings.ToLower(*req.Role))
	if userRole == entity.UserRoleAdmin {
		return fmt.Errorf("user role '%s' cannot be created", *req.Role)
	}

	if !utils.ValueInSlice(userRole, []entity.UserRole{entity.UserRoleEditor, entity.UserRoleReporter, entity.UserRoleReader}) {
		return fmt.Errorf("invalid user role '%s'", *req.Role)
	}

	return nil
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

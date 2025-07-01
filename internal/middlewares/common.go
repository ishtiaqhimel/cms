package middlewares

import (
	"context"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

// RoleBasedAccessControl is the middleware func that validates user role-based access control logic
func RoleBasedAccessControl(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleKey := c.Request().Header.Get(utils.HeaderKeyUserRoles)
		if roleKey == "" {
			return c.JSON(response.RespondError(response.ErrBadRequest,
				fmt.Errorf("user role header %s is empty", utils.HeaderKeyUserRoles)))
		}

		userRoles := strings.Split(roleKey, ",")
		for i := range userRoles {
			userRoles[i] = strings.TrimSpace(userRoles[i])
		}

		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), utils.ContextKeyUserRoles, userRoles)))

		return h(c)
	}
}

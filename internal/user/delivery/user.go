package delivery

import (
	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/user/usecase"
)

// userHandler represents the httpHandler for user
type userHandler struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(e *echo.Echo, usecase usecase.UserUsecase) {
	handler := &userHandler{
		Usecase: usecase,
	}

	userV1 := e.Group("/api/v1")

	userV1.POST("/user", handler.CreateUser)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	req := new(model.UserCreateReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	userID, err := h.Usecase.CreateUser(req)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}
	return c.JSON(response.RespondCreated("request is successful", userID))
}

package delivery

import (
	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
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

	userV1 := e.Group("/api/v1/user")

	userV1.POST("", handler.CreateUser)
	userV1.GET("/:user_id", handler.GetUserByID)
	userV1.PUT("/:user_id", handler.UpdateUser)
	userV1.GET("", handler.ListUserByFilter)
	userV1.DELETE("/:user_id", handler.DeleteUserByID)
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

func (h *userHandler) GetUserByID(c echo.Context) error {
	userID := c.Param("user_id")

	user, err := h.Usecase.GetUserByID(userID)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccess("request is successful", user))
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	userID := c.Param("user_id")

	req := new(model.UserUpdateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := h.Usecase.UpdateUser(userID, req); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

func (h *userHandler) ListUserByFilter(c echo.Context) error {
	req := &usecase.UserListFilter{}
	if err := utils.RequestQueryParamToStruct(c, []string{
		"id", "name", "email", "role",
	}, &req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	page, pageSize, err := utils.PaginationParams(c, config.Get().App.DefaultPageSize, config.Get().App.MaxPageSize)
	if err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	res, total, err := h.Usecase.ListUserByFilter(req, &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessForList("request is successful", int(total), pageSize, page, res))
}

func (h *userHandler) DeleteUserByID(c echo.Context) error {
	userID := c.Param("user_id")

	if err := h.Usecase.DeleteUserByID(userID); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

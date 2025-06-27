package delivery

import (
	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
)

// categoryHandler represents the httpHandler for category
type categoryHandler struct {
	Usecase usecase.CategoryUsecase
}

func NewCategoryHandler(e *echo.Echo, usecase usecase.CategoryUsecase) {
	handler := &categoryHandler{
		Usecase: usecase,
	}

	categoryV1 := e.Group("/api/v1")

	categoryV1.POST("/category", handler.CreateCategory)
}

func (h *categoryHandler) CreateCategory(c echo.Context) error {
	req := new(model.CategoryCreateReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := defaults.Set(req); err != nil {
		return c.JSON(response.RespondError(err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	categoryID, err := h.Usecase.CreateCategory(req)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}
	return c.JSON(response.RespondCreated("request is successful", categoryID))
}

package delivery

import (
	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/category/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/category/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

// categoryHandler represents the httpHandler for category
type categoryHandler struct {
	Usecase usecase.CategoryUsecase
}

func NewCategoryHandler(e *echo.Echo, usecase usecase.CategoryUsecase) {
	handler := &categoryHandler{
		Usecase: usecase,
	}

	categoryV1 := e.Group("/api/v1/category")

	categoryV1.POST("", handler.CreateCategory)
	categoryV1.GET("/:category_id", handler.GetCategoryByID)
	categoryV1.PUT("/:category_id", handler.UpdateCategory)
	categoryV1.GET("", handler.ListCategoryByFilter)
	categoryV1.DELETE("/:category_id", handler.DeleteCategoryByID)
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

func (h *categoryHandler) GetCategoryByID(c echo.Context) error {
	categoryID := c.Param("category_id")

	category, err := h.Usecase.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccess("request is successful", category))
}

func (h *categoryHandler) UpdateCategory(c echo.Context) error {
	categoryID := c.Param("category_id")

	req := new(model.CategoryUpdateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := h.Usecase.UpdateCategory(categoryID, req); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

func (h *categoryHandler) ListCategoryByFilter(c echo.Context) error {
	req := &usecase.CategoryListFilter{}
	if err := utils.RequestQueryParamToStruct(c, []string{
		"id", "name", "parent_id", "is_active",
	}, &req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	page, pageSize, err := utils.PaginationParams(c, config.Get().App.DefaultPageSize, config.Get().App.MaxPageSize)
	if err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	res, total, err := h.Usecase.ListCategoryByFilter(req, &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessForList("request is successful", int(total), pageSize, page, res))
}

func (h *categoryHandler) DeleteCategoryByID(c echo.Context) error {
	categoryID := c.Param("category_id")

	if err := h.Usecase.DeleteCategoryByID(categoryID); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

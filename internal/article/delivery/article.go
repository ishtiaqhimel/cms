package delivery

import (
	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/ishtiaqhimel/news-portal/cms/internal/middlewares"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

// articleHandler represents the httpHandler for article
type articleHandler struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, usecase usecase.ArticleUsecase) {
	handler := &articleHandler{
		Usecase: usecase,
	}

	articleV1 := e.Group("/api/v1/article")

	articleV1.POST("", handler.CreateArticle, middlewares.RoleBasedAccessControl)
	articleV1.GET("/:article_id", handler.GetArticleByID)
	articleV1.PUT("/:article_id", handler.UpdateArticle)
	articleV1.GET("", handler.ListArticleByFilter)
	articleV1.DELETE("/:article_id", handler.DeleteArticleByID)
}

func (h *articleHandler) CreateArticle(c echo.Context) error {
	req := new(model.ArticleCreateReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	articleID, err := h.Usecase.CreateArticle(c.Request().Context(), req)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}
	return c.JSON(response.RespondCreated("request is successful", articleID))
}

func (h *articleHandler) GetArticleByID(c echo.Context) error {
	articleID := c.Param("article_id")

	article, err := h.Usecase.GetArticleByID(articleID)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccess("request is successful", article))
}

func (h *articleHandler) UpdateArticle(c echo.Context) error {
	articleID := c.Param("article_id")

	req := new(model.ArticleUpdateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := h.Usecase.UpdateArticle(articleID, req); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

func (h *articleHandler) ListArticleByFilter(c echo.Context) error {
	req := &usecase.ArticleListFilter{}
	if err := utils.RequestQueryParamToStruct(c, []string{
		"id", "status", "category_id", "author_id",
	}, &req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	page, pageSize, err := utils.PaginationParams(c, config.Get().App.DefaultPageSize, config.Get().App.MaxPageSize)
	if err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	res, total, err := h.Usecase.ListArticleByFilter(req, &utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessForList("request is successful", int(total), pageSize, page, res))
}

func (h *articleHandler) DeleteArticleByID(c echo.Context) error {
	articleID := c.Param("article_id")

	if err := h.Usecase.DeleteArticleByID(articleID); err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccessWithNoContent("request is successful"))
}

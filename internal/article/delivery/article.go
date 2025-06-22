package delivery

import (
	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
)

// ArticleHandler represents the httpHandler for article
type ArticleHandler struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, usecase usecase.ArticleUsecase) {
	handler := &ArticleHandler{
		Usecase: usecase,
	}

	articleV1 := e.Group("/api/v1")

	articleV1.POST("/article", handler.CreateArticle)
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	req := new(model.ArticleCreateReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(response.RespondError(response.ErrBadRequest, err))
	}

	articleID, err := h.Usecase.CreateArticle(req)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}
	return c.JSON(response.RespondCreated("request is successful", articleID))
}

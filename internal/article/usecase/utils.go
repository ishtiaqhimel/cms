package usecase

import (
	"github.com/ishtiaqhimel/news-portal/cms/entity"
	"github.com/ishtiaqhimel/news-portal/cms/internal/article/model"
)

type ArticleListFilter struct {
	ID         *string
	Status     *string
	CategoryID *string
	AuthorID   *string
}

func articleCreateReqToArticle(req *model.ArticleCreateReq) *entity.Article {
	return &entity.Article{
		Title:      req.Title,
		Body:       req.Body,
		Status:     entity.ArticleStatusDraft,
		CategoryID: req.CategoryID,
		AuthorID:   req.AuthorID,
	}
}

func articleUpdateReqToArticle(id string, req *model.ArticleUpdateReq) (*entity.Article, []string) {
	article := &entity.Article{
		ID: id,
	}
	var updateCols []string

	if req.Title != nil {
		article.Title = *req.Title
		updateCols = append(updateCols, "title")
	}

	if req.Body != nil {
		article.Body = *req.Body
		updateCols = append(updateCols, "body")
	}

	if req.CategoryID != nil {
		article.CategoryID = *req.CategoryID
		updateCols = append(updateCols, "category_id")
	}

	if req.AuthorID != nil {
		article.AuthorID = *req.AuthorID
		updateCols = append(updateCols, "author_id")
	}

	return article, updateCols
}

func toArticleResp(article *entity.Article) *model.ArticleResp {
	if article == nil {
		return nil
	}

	return &model.ArticleResp{
		ID:         article.ID,
		Title:      article.Title,
		Body:       article.Body,
		Status:     string(article.Status),
		CategoryID: article.CategoryID,
		AuthorID:   article.AuthorID,
	}
}

func toArticleRespList(in []*entity.Article) []*model.ArticleResp {
	out := make([]*model.ArticleResp, 0)

	for _, i := range in {
		temp := toArticleResp(i)
		if temp != nil {
			out = append(out, temp)
		}
	}

	return out
}

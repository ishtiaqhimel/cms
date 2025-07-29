package model

// Request Models

type ArticleCreateReq struct {
	Title      string `json:"title" validate:"required"`
	Body       string `json:"body" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	AuthorID   string `json:"author_id" validate:"required"`
}

type ArticleUpdateReq struct {
	Title      *string `json:"title"`
	Body       *string `json:"body"`
	CategoryID *string `json:"category_id"`
	AuthorID   *string `json:"author_id"`
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

type ArticleResp struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Status     string `json:"status"`
	CategoryID string `json:"category_id"`
	AuthorID   string `json:"author_id"`
}

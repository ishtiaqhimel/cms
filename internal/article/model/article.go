package model

// Request Models

type ArticleCreateReq struct {
	Title      string `json:"title" validate:"required"`
	Body       string `json:"body" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	AuthorID   string `json:"author_id" validate:"required"`
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

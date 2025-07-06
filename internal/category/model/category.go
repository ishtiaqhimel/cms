package model

// Request Models

type CategoryCreateReq struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	IsActive    bool    `json:"is_active" default:"true"`
}

type CategoryUpdateReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	IsActive    *bool   `json:"is_active"`
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

type CategoryResp struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
	IsActive    bool    `json:"is_active"`
}

package model

// Request Models

type CategoryCreateReq struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	IsActive    bool    `json:"is_active" default:"true"`
}

// Response Models

type EntityID struct {
	ID string `json:"id"`
}

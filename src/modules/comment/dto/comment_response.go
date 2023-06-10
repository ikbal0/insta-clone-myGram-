package dto

import "time"

type CommentResponse struct {
	Message   string     `json:"message" form:"message"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

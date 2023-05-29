package dto

type CommentRequestBody struct {
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" form:"message"`
}

package dto

type PhotoRequestBody struct {
	Title   uint `json:"title" form:"title"`
	Caption uint `json:"caption" form:"caption"`
}

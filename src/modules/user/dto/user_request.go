package dto

type UserRequestBody struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

type UserLoginRequestBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

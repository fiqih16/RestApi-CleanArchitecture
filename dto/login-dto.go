package dto

// LoginDto is a model that used by client when POST from /login url
type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
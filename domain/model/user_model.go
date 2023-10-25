package model

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

package models

import "time"

type SwaggerUserRegister struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type SwaggerLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwaggerSuccessLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SwaggerRegisterInput struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SwaggerSuccessRegisterResponse struct {
	Message string              `json:"message"`
	User    SwaggerUserRegister `json:"user"`
}

type SwaggerTodo struct {
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
}

type SwaggerTodoCreate struct {
	Title string `json:"title"`
}

type SwaggerTodoCreateResponse struct {
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}

type SwaggerErrorResponse struct {
	Error string `json:"error"`
}

type SwaggerSuccessResponse struct {
	Message string `json:"message"`
}

type SwaggerToggleSuccessResponse struct {
	Message string      `json:"message"`
	Todo    SwaggerTodo `json:"todo"`
}

type SwaggerGetTodosResponse struct {
	Message string        `json:"message"`
	Todo    []SwaggerTodo `json:"todo"`
}

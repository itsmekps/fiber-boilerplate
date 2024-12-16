package dtos

type GetUserDetailsRequest struct {}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"omitempty,min=3,max=20"`
	Email    string `json:"email" validate:"omitempty,email"`
}

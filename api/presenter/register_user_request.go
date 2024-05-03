package presenter

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Email    string `json:"email" validate:"required,min=5,max=255,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

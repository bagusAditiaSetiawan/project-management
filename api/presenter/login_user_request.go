package presenter

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

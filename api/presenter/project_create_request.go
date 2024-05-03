package presenter

type ProjectCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsActive    bool   `json:"is_active" validate:"required,boolean"`
}

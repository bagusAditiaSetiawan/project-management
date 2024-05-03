package presenter

type PaginationRequest struct {
	Page  int `json:"page" validate:"required,number,min=1"`
	Limit int `json:"limit" validate:"required,number,min=1"`
}

func (pagination *PaginationRequest) GetOffset() int {
	return (pagination.Page - 1) * pagination.Limit
}

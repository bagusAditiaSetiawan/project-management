package presenter

type ProjectPaginationRequest struct {
	PaginationRequest
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

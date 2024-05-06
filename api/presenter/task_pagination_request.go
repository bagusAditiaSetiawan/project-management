package presenter

type TaskPaginationRequest struct {
	PaginationRequest
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

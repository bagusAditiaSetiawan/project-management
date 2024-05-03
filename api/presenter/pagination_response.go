package presenter

type PaginationResponse struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

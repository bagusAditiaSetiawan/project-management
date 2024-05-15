package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"net/http"
)

func ToResponsePagination(response presenter.PaginationResponse) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
}

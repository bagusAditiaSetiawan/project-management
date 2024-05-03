package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"net/http"
)

func ToResponseProjectPagination(response presenter.PaginationResponse) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
}

package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"net/http"
)

func ToReadResponse(data interface{}) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}
}

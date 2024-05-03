package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"net/http"
)

func ToCreatedResponse(data interface{}) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   data,
	}
}

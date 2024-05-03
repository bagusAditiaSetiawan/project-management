package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"net/http"
)

func ToResponseBadRequest(message string) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   message,
	}
}

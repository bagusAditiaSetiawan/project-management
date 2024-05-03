package exception

import (
	"errors"
	"net/http"

	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandlerException(ctx *fiber.Ctx, err error) error {
	if errNotFound := ErrorNotFoundHandler(err); errNotFound != nil {
		return ctx.Status(http.StatusNotFound).JSON(&presenter.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   errNotFound.Error(),
		})
	}

	if errValidation := ErrorValidator(err); errValidation != nil {
		customValidator := helpers.CustomErrorValidator(errValidation)
		return ctx.Status(http.StatusBadRequest).JSON(&presenter.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   customValidator,
		})
	}

	if errBadRequestException := BadRequestExceptionHandler(err); errBadRequestException != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&presenter.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   errBadRequestException.Error(),
		})
	}

	if errorBodyException := ErrorBodyExceptionHandler(err); errorBodyException != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&presenter.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   errorBodyException.Error(),
		})
	}

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	if err != nil {
		return ctx.Status(code).JSON(&presenter.WebResponse{
			Code:   code,
			Status: "",
			Data:   err.Error(),
		})
	}

	return nil
}

func ErrorValidator(err error) validator.ValidationErrors {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		return exception
	} else {
		return nil
	}
}

func ErrorNotFoundHandler(err error) error {
	exception, ok := err.(NotFoundHandler)
	if ok {
		return exception
	} else {
		return nil
	}
}
func BadRequestExceptionHandler(err error) error {
	exception, ok := err.(BadRequestException)
	if ok {
		return exception
	} else {
		return nil
	}
}
func ErrorBodyExceptionHandler(err error) error {
	exception, ok := err.(ErrorBodyException)
	if ok {
		return exception
	} else {
		return nil
	}
}

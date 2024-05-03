package handler

import (
	"fmt"
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthControllerImpl struct {
	UserService auth.UserService
}

func NewAuthControllerImpl(userService auth.UserService) *AuthControllerImpl {
	return &AuthControllerImpl{
		UserService: userService,
	}
}

func (controller *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	registerUserRequest := new(presenter.RegisterUserRequest)
	err := ctx.BodyParser(registerUserRequest)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	register := controller.UserService.Register(registerUserRequest)

	return ctx.JSON(helpers.ToResponseRegisterResponse(register))
}

func (controller *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	loginRequest := new(presenter.LoginUserRequest)
	err := ctx.BodyParser(loginRequest)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	accessToken := controller.UserService.Login(loginRequest)

	return ctx.JSON(helpers.ToResponseLoginResponse(accessToken))
}

func (controller *AuthControllerImpl) Profile(ctx *fiber.Ctx) error {
	token := ctx.Locals("auth").(*jwt.Token)
	fmt.Println(ctx.Locals("auth"), token)
	return nil
}

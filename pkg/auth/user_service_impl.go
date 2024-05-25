package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB              *gorm.DB
	UserRepository  UserRepository
	Validator       *validator.Validate
	PasswordService PasswordService
	JwtService      JwtService
	Logger          *aws_cloudwatch.AwsCloudWatchServiceImpl
}

func NewUserServiceImpl(db *gorm.DB, userRepository UserRepository, validator *validator.Validate, passwordService PasswordService, jwtService JwtService, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) *UserServiceImpl {
	return &UserServiceImpl{
		DB:              db,
		UserRepository:  userRepository,
		Validator:       validator,
		PasswordService: passwordService,
		JwtService:      jwtService,
		Logger:          logger,
	}
}

func (service *UserServiceImpl) Register(request *presenter.RegisterUserRequest) entities.User {
	err := service.Validator.Struct(request)
	helpers.IfPanicHelper(err)

	hashedPassword, err := service.PasswordService.Hashing(request.Password)
	helpers.IfPanicHelper(err)
	request.Password = hashedPassword

	tx := service.DB.Begin()
	user, err := service.UserRepository.Register(tx, request)
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}

	defer helpers.RollbackOrCommitDb(tx)

	return user
}

func (service *UserServiceImpl) Login(request *presenter.LoginUserRequest) string {
	err := service.Validator.Struct(request)
	helpers.IfPanicHelper(err)
	tx := service.DB.Begin()

	go service.Logger.SendLogInfo("Process login user", request.Username)

	user := service.UserRepository.FindByEmailUsername(tx, "", request.Username)
	if user.ID == 0 {
		panic(exception.NewNotFoundHandler("User is not found"))
	}
	isMatched := service.PasswordService.Compare(request.Password, user.Password)
	if !isMatched {
		panic(exception.NewBadRequestException("Password is wrong"))
	}

	defer helpers.RollbackOrCommitDb(tx)

	accessToken := service.JwtService.Sign(user)
	return accessToken
}

func (service *UserServiceImpl) ClaimsUserToken(token *jwt.Token) entities.User {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(int)
	tx := service.DB.Begin()
	user := service.UserRepository.FindById(tx, userId)
	defer helpers.RollbackOrCommitDb(tx)
	return user
}

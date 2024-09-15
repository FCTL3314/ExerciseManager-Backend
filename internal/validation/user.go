package validation

import (
	"ExerciseManager/internal/domain"
	"github.com/go-playground/validator/v10"
)

type IUserValidator interface {
	ValidateCreateUserRequest(createUserRequest *domain.CreateUserRequest) error
	ValidateLoginUserRequest(loginUserRequest *domain.LoginUserRequest) error
	ValidateRefreshTokenRequest(refreshTokenRequest *domain.RefreshTokenRequest) error
	ValidateUpdateUserRequest(updateUserRequest *domain.UpdateUserRequest) error
}

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		validate: validator.New(),
	}
}

func (uv *UserValidator) ValidateCreateUserRequest(createUserRequest *domain.CreateUserRequest) error {
	return uv.validate.Struct(createUserRequest)
}

func (uv *UserValidator) ValidateLoginUserRequest(loginUserRequest *domain.LoginUserRequest) error {
	return uv.validate.Struct(loginUserRequest)
}

func (uv *UserValidator) ValidateRefreshTokenRequest(refreshTokenRequest *domain.RefreshTokenRequest) error {
	return uv.validate.Struct(refreshTokenRequest)
}

func (uv *UserValidator) ValidateUpdateUserRequest(updateUserRequest *domain.UpdateUserRequest) error {
	return uv.validate.Struct(updateUserRequest)
}

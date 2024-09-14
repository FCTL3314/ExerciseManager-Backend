package validation

import (
	"ExerciseManager/internal/domain"
	"github.com/go-playground/validator/v10"
)

type IUserValidator interface {
	ValidateCreateUser(createUser *domain.CreateUser) error
	ValidateLoginUser(createUser *domain.LoginUser) error
	ValidateUpdateUser(updateUser *domain.UpdateUser) error
}

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		validate: validator.New(),
	}
}

func (uv *UserValidator) ValidateCreateUser(createUser *domain.CreateUser) error {
	return uv.validate.Struct(createUser)
}

func (uv *UserValidator) ValidateLoginUser(createUser *domain.LoginUser) error {
	return uv.validate.Struct(createUser)
}

func (uv *UserValidator) ValidateUpdateUser(updateUser *domain.UpdateUser) error {
	return uv.validate.Struct(updateUser)
}

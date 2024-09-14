package validation

import (
	"ExerciseManager/internal/domain"
	"github.com/go-playground/validator/v10"
)

type UserValidator interface {
	ValidateCreateUser(createUser *domain.CreateUser) error
	ValidateUpdateUser(updateUser *domain.UpdateUser) error
}

type DefaultUserValidator struct {
	validate *validator.Validate
}

func NewDefaultUserValidator() *DefaultUserValidator {
	return &DefaultUserValidator{
		validate: validator.New(),
	}
}

func (uv *DefaultUserValidator) ValidateCreateUser(createUser *domain.CreateUser) error {
	return uv.validate.Struct(createUser)
}

func (uv *DefaultUserValidator) ValidateUpdateUser(updateUser *domain.UpdateUser) error {
	return uv.validate.Struct(updateUser)
}

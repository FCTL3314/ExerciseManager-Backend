package bootstrap

import (
	"ExerciseManager/internal/validation"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate, cfg *Config) error {
	if err := registerImageExtensionValidator(v, cfg.Uploads.AllowedImageExtensions); err != nil {
		return err
	}
	return nil
}

func registerImageExtensionValidator(v *validator.Validate, AllowedImageExtensions []string) error {
	if err := v.RegisterValidation(
		"imageextension",
		func(fl validator.FieldLevel) bool {
			return validation.IsValidImageExtension(
				fl.Field().String(),
				AllowedImageExtensions,
			)
		},
	); err != nil {
		return err
	}
	return nil
}

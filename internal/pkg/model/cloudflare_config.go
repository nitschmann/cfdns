package model

import (
	"github.com/go-playground/validator/v10"
)

type CloudflareConfig struct {
	ApiKey string `validate:"required"`
	Email  string `validate:"required,email"`
}

func (c CloudflareConfig) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

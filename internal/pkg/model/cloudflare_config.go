package model

import (
	"github.com/go-playground/validator/v10"
)

// CloudflareConfig is the data struct for the API configuration for Cloudflare
type CloudflareConfig struct {
	APIKey string `validate:"required"`
	Email  string `validate:"required,email"`
}

// Validate validates the contents of the CloudflareConfig struct
func (c CloudflareConfig) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

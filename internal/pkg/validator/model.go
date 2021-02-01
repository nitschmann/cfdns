package validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ModelValidator is the interface to validate model objects
type ModelValidator interface {
	IsValid() bool
	Validate() error
}

// ModelValidatorObj implements the ModelValidator per default
type ModelValidatorObj struct {
	Obj       model.Model
	Validator *validator.Validate
}

// NewModelValidator creates an new pointer instance of ModelValidatorObj with default values
func NewModelValidator(obj model.Model) *ModelValidatorObj {
	return &ModelValidatorObj{
		Obj:       obj,
		Validator: validator.New(),
	}
}

// IsValid check if v.Validate() returns any error or not
func (v *ModelValidatorObj) IsValid() bool {
	err := v.Validate()
	return err == nil
}

// Validate runs the validations in the v.Obj (if defined)
// TODO: Add custom error message here
func (v *ModelValidatorObj) Validate() error {
	err := v.Validator.Struct(v.Obj)
	return err
}

package validator

import val "github.com/go-playground/validator/v10"

// custom echo validator
type validator struct {
	validator *val.Validate
}

func New() *validator {
	return &validator{validator: val.New()}
}

// validates a value
func (v *validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

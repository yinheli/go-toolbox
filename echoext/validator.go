package echoext

import (
	"reflect"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func ConfigEchoValidator(e *echo.Echo) {
	validator := NewValidator()
	e.Validator = validator
	e.Binder = validator.WrapBinder(e.Binder)
}

type Validator struct {
	originBinder echo.Binder
}

func NewValidator() *Validator {
	return &Validator{}
}

type ValidationError struct {
	msg string
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{msg}
}

func (v *ValidationError) Error() string {
	return v.msg
}

func (v *Validator) Validate(i interface{}) error {
	valid := validate.New(i)
	if valid.Validate() {
		return nil
	}
	return NewValidationError(valid.Errors.One())
}

func (v *Validator) WrapBinder(originBinder echo.Binder) echo.Binder {
	v.originBinder = originBinder
	return v
}

// validate for deep validate
func (v *Validator) validate(i interface{}) error {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Ptr, reflect.Struct:
			if field.CanInterface() {
				if err := v.validate(field.Interface()); err != nil {
					return err
				}
			}
		}
	}

	return v.Validate(i)
}

func (v *Validator) Bind(i interface{}, c echo.Context) error {
	err := v.originBinder.Bind(i, c)
	if err == nil {
		return v.validate(i)
	}
	return err
}

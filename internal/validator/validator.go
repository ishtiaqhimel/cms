package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(i interface{}) error
}

type genericValidator struct {
	Validator *validator.Validate
}

func New() Validator {
	v := &genericValidator{
		Validator: validator.New(),
	}

	// register a function to get json tags instead of go struct field names
	v.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

func (v *genericValidator) Validate(s interface{}) error {
	if err := v.Validator.Struct(s); err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("invalid '%s' field", err.Field())
		}
	}
	return nil
}

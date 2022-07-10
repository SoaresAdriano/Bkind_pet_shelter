package domain

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidatorInterface determines the contracts with validator
type ValidatorInterface interface {
	Validate(obj interface{}) (err []error)
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator(ctx context.Context) (Validator, error) {
	return Validator{}, ErrNotImplemented
}

func MustNewValidator(ctx context.Context) *Validator {
	validator, err := NewValidator(ctx)
	if err != nil {
		panic(err)
	}
	return &validator
}

func (v Validator) Validate(obj interface{}) error {
	err := v.validator.Struct(obj)
	if err == nil {
		return err
	}

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Printf("field validation error: %s \n tag: %s", e.Field(), e.Tag())
	}
	return err
}

package domain

import (
	"context"

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

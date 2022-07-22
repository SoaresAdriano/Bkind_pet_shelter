package domain

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type Logger interface {
	Debug(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})
}

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

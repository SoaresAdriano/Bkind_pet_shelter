package domain

import "errors"

var (
	// ErrNotImplemented indicates that the operation or functionality is not implemented yet
	ErrNotImplemented = errors.New("not implemented")

	// ErrMissingPortDependency indicates that was not provided the port param
	ErrMissingPortDependency = errors.New("missing required dependency: port")

	// ErrMissingAppDependency indicates that was not provided the app param
	ErrMissingAppDependency = errors.New("missing required dependency: app")
)

package errors

import "errors"

var (
	// ErrNotFound indicates that a requested resource was not found.
	ErrNotFound = errors.New("not found")

	// ErrPermissionDenied indicates that the caller does not have permission to perform the action.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrInvalidArgument indicates that the client specified an invalid argument.
	ErrInvalidArgument = errors.New("invalid argument")
)
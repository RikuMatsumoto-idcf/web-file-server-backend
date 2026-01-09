package domain

import "errors"

// Domain errors represent business rule violations
// These errors should be mapped to appropriate HTTP status codes in the handler layer

var (
	// ErrInvalidName is returned when a file name is invalid
	// Example: empty name, contains path traversal characters, etc.
	ErrInvalidName = errors.New("invalid file name")

	// ErrNotFound is returned when a requested file does not exist
	ErrNotFound = errors.New("file not found")

	// ErrAlreadyExists is returned when trying to create a file that already exists
	ErrAlreadyExists = errors.New("file already exists")

	// ErrTooLarge is returned when a file exceeds the maximum allowed size
	ErrTooLarge = errors.New("file too large")
)

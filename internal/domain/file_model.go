package domain

import "time"

// File represents a file entity in the domain layer
// This is the core business object that represents a file
type File struct {
	ID        int       // Unique identifier
	Name      string    // File name
	Data      []byte    // File content (binary data)
	CreatedAt time.Time // Creation timestamp
	UpdatedAt time.Time // Last update timestamp
}

// FileName is a value object for validated file names
// Use NewFileName to create validated instances
type FileName string

// NewFileName creates a validated FileName
// TODO: Implement validation logic
// - Check for empty strings
// - Prevent path traversal (reject "..", "/", "\\")
// - Apply any business rules for file naming
func NewFileName(raw string) (FileName, error) {
	panic("implement me")
}

// String returns the string representation of FileName
func (f FileName) String() string {
	return string(f)
}

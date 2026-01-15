package usecase

import (
	"context"
	"io"

	"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/domain"
)

// FileUsecase defines the business logic interface for file operations
// This interface represents the application's use cases
type FileUsecase interface {
	// Upload stores a new file
	// Returns domain.ErrAlreadyExists if file already exists
	// Returns domain.ErrInvalidName if name validation fails
	// Returns domain.ErrTooLarge if file size exceeds limits
	Upload(ctx context.Context, name string, data io.Reader) error

	// Download retrieves a file's content
	// Returns domain.ErrNotFound if file doesn't exist
	// Returns domain.ErrInvalidName if name validation fails
	// The caller is responsible for closing the returned io.ReadCloser
	Download(ctx context.Context, name string) (io.ReadCloser, error)

	// List returns all file names
	// Can be used to browse available files
	List(ctx context.Context) ([]string, error)

	// Delete removes a file
	// Returns domain.ErrNotFound if file doesn't exist
	// Returns domain.ErrInvalidName if name validation fails
	Delete(ctx context.Context, name string) error
}

// FileRepository defines the interface for file persistence
// This interface is implemented by the infrastructure layer
type FileRepository interface {
	// Save stores file data
	Save(ctx context.Context, name domain.FileName, data io.Reader) error

	// FindByName retrieves file data
	// Returns domain.ErrNotFound if not found
	FindByName(ctx context.Context, name domain.FileName) (io.ReadCloser, error)

	// Exists checks if a file exists
	Exists(ctx context.Context, name domain.FileName) (bool, error)

	// ListAll returns all file names
	ListAll(ctx context.Context) ([]domain.FileName, error)

	// Remove deletes a file
	Remove(ctx context.Context, name domain.FileName) error
}

// fileUsecaseImpl is the concrete implementation of FileUsecase
type fileUsecaseImpl struct {
	repo FileRepository
}

// NewFileUsecase creates a new FileUsecase instance
// repo: the repository implementation for data persistence
func NewFileUsecase(repo FileRepository) FileUsecase {
	return &fileUsecaseImpl{
		repo: repo,
	}
}

// Upload implements FileUsecase.Upload
// TODO: Implement the following logic:
// 1. Validate the file name using domain.NewFileName
// 2. Check if file already exists using repo.Exists
// 3. If exists, return domain.ErrAlreadyExists
// 4. Save the file using repo.Save
// 5. Handle any errors appropriately
func (u *fileUsecaseImpl) Upload(ctx context.Context, name string, data io.Reader) error {
	panic("implement me")
}

// Download implements FileUsecase.Download
// TODO: Implement the following logic:
// 1. Validate the file name using domain.NewFileName
// 2. Retrieve file data using repo.FindByName
// 3. Return the io.ReadCloser (caller must close it)
// 4. Handle domain.ErrNotFound appropriately
func (u *fileUsecaseImpl) Download(ctx context.Context, name string) (io.ReadCloser, error) {
	panic("implement me")
}

// List implements FileUsecase.List
// TODO: Implement the following logic:
// 1. Call repo.ListAll to get all file names
// 2. Convert domain.FileName slice to string slice
// 3. Return the list
func (u *fileUsecaseImpl) List(ctx context.Context) ([]string, error) {
	panic("implement me")
}

// Delete implements FileUsecase.Delete
// TODO: Implement the following logic:
// 1. Validate the file name using domain.NewFileName
// 2. Check if file exists using repo.Exists
// 3. If not exists, return domain.ErrNotFound
// 4. Remove the file using repo.Remove
func (u *fileUsecaseImpl) Delete(ctx context.Context, name string) error {
	panic("implement me")
}

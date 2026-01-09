package infrastructure

import (
"context"
"io"

"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/domain"
"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/usecase"
)

// fileRepository implements usecase.FileRepository interface
// This is the infrastructure layer that handles actual file storage
type fileRepository struct {
// You can add fields here for database connection, file system path, etc.
// Example: db *sql.DB, storagePath string
}

// NewFileRepository creates a new FileRepository instance
// TODO: Add necessary dependencies (e.g., database connection, storage path)
func NewFileRepository( /* add dependencies here */ ) usecase.FileRepository {
return &fileRepository{
// Initialize fields here
}
}

// Save implements usecase.FileRepository.Save
// TODO: Implement file storage logic:
// 1. Read data from io.Reader
// 2. Store to database or file system
// 3. Handle duplicate file names appropriately
// 4. Return errors if storage fails
func (r *fileRepository) Save(ctx context.Context, name domain.FileName, data io.Reader) error {
panic("implement me")
}

// FindByName implements usecase.FileRepository.FindByName
// TODO: Implement file retrieval logic:
// 1. Query database or file system for the file
// 2. If not found, return domain.ErrNotFound
// 3. Return an io.ReadCloser containing the file data
// 4. Ensure the returned ReadCloser can be closed by the caller
func (r *fileRepository) FindByName(ctx context.Context, name domain.FileName) (io.ReadCloser, error) {
panic("implement me")
}

// Exists implements usecase.FileRepository.Exists
// TODO: Implement existence check:
// 1. Check if file exists in database or file system
// 2. Return true if exists, false otherwise
// 3. Handle errors appropriately
func (r *fileRepository) Exists(ctx context.Context, name domain.FileName) (bool, error) {
panic("implement me")
}

// ListAll implements usecase.FileRepository.ListAll
// TODO: Implement listing logic:
// 1. Query database or scan file system
// 2. Return slice of all file names
// 3. Handle empty list case (return empty slice, not error)
func (r *fileRepository) ListAll(ctx context.Context) ([]domain.FileName, error) {
panic("implement me")
}

// Remove implements usecase.FileRepository.Remove
// TODO: Implement deletion logic:
// 1. Delete file from database or file system
// 2. If file doesn't exist, return domain.ErrNotFound
// 3. Handle deletion errors
func (r *fileRepository) Remove(ctx context.Context, name domain.FileName) error {
panic("implement me")
}

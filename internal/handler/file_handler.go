package handler

import (
"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/usecase"
"github.com/labstack/echo/v4"
)

// FileHandler handles HTTP requests for file operations
// This is the presentation layer that interacts with Echo framework
type FileHandler struct {
fileUsecase usecase.FileUsecase
}

// NewFileHandler creates a new FileHandler instance
// fileUsecase: the business logic layer for file operations
func NewFileHandler(fileUsecase usecase.FileUsecase) *FileHandler {
return &FileHandler{
fileUsecase: fileUsecase,
}
}

// Upload handles file upload requests
// Route: PUT /api/files/:name
// TODO: Implement the following:
// 1. Get the file name from URL parameter using c.Param("name")
// 2. Get request body using c.Request().Body
// 3. Call fileUsecase.Upload with context, name, and body
// 4. Map domain errors to HTTP status codes:
//    - domain.ErrInvalidName -> 400 Bad Request
//    - domain.ErrAlreadyExists -> 409 Conflict
//    - domain.ErrTooLarge -> 413 Payload Too Large
// 5. Return 204 No Content on success
// 6. Use c.JSON() for error responses
func (h *FileHandler) Upload(c echo.Context) error {
panic("implement me")
}

// Download handles file download requests
// Route: GET /api/files/:name
// TODO: Implement the following:
// 1. Get the file name from URL parameter using c.Param("name")
// 2. Call fileUsecase.Download with context and name
// 3. Map domain errors to HTTP status codes:
//    - domain.ErrNotFound -> 404 Not Found
//    - domain.ErrInvalidName -> 400 Bad Request
// 4. Set Content-Type header to "application/octet-stream"
// 5. Stream the file content to response using c.Stream()
// 6. Don't forget to close the ReadCloser (use defer)
func (h *FileHandler) Download(c echo.Context) error {
panic("implement me")
}

// List handles file listing requests
// Route: GET /api/files
// TODO: Implement the following:
// 1. Call fileUsecase.List with context
// 2. Return JSON array of file names using c.JSON(200, fileNames)
// 3. Handle errors appropriately
func (h *FileHandler) List(c echo.Context) error {
panic("implement me")
}

// Delete handles file deletion requests
// Route: DELETE /api/files/:name
// TODO: Implement the following:
// 1. Get the file name from URL parameter using c.Param("name")
// 2. Call fileUsecase.Delete with context and name
// 3. Map domain errors to HTTP status codes:
//    - domain.ErrNotFound -> 404 Not Found
//    - domain.ErrInvalidName -> 400 Bad Request
// 4. Return 204 No Content on success
func (h *FileHandler) Delete(c echo.Context) error {
panic("implement me")
}

// RegisterRoutes registers all file-related routes to the Echo instance
// This method sets up the routing for file operations
func (h *FileHandler) RegisterRoutes(e *echo.Echo) {
// File operations group
files := e.Group("/api/files")

// Route definitions
files.GET("", h.List)              // List all files
files.PUT("/:name", h.Upload)      // Upload a file
files.GET("/:name", h.Download)    // Download a file
files.DELETE("/:name", h.Delete)   // Delete a file
}

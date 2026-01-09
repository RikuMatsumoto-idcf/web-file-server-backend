package main

import (
	"log"

	"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/handler"
	"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/infrastructure"
	"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logging middleware
	e.Use(middleware.Recover()) // Recover from panics
	e.Use(middleware.CORS())    // CORS middleware

	// Dependency Injection (DI)
	// TODO: Initialize dependencies properly
	// 1. Set up database connection or file storage
	// 2. Create repository instance with actual dependencies
	// Example:
	//   db := setupDatabase()
	//   fileRepo := infrastructure.NewFileRepository(db)
	
	// For now, this will panic when called because NewFileRepository needs parameters
	// You need to implement the repository initialization first
	fileRepo := infrastructure.NewFileRepository()
	fileUsecase := usecase.NewFileUsecase(fileRepo)
	fileHandler := handler.NewFileHandler(fileUsecase)

	// Register routes
	fileHandler.RegisterRoutes(e)

	// Start server
	port := ":8080"
	log.Printf("Server is starting on http://localhost%s\n", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

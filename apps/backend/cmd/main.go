package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/raihanstark/pomodeep/internal/database"
	"github.com/raihanstark/pomodeep/internal/usecase/users"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Database configuration
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "pomodeep"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Connect to database
	pool, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// Initialize dependencies
	// !LEARN: Dependency Injection
	queries := database.NewQueries(pool)
	userRepo := users.NewRepository(queries)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "healthy",
			"message": "PomoDeep API is running",
		})
	})

	// API routes
	api := e.Group("/api/v1")
	api.POST("/auth/signup", userHandler.SignUp)
	api.POST("/auth/signin", userHandler.SignIn)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

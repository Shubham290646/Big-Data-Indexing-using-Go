package main

import (
	"rest-api-json-handler/handlers"
	"rest-api-json-handler/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Redis
	storage.InitRedis()

	// Create a new Echo instance
	e := echo.New()

	// Define API routes
	e.POST("/plans", handlers.CreatePlan)       // Create a plan
	e.GET("/plans/:id", handlers.GetPlan)       // Retrieve a plan (Conditional Read)
	e.DELETE("/plans/:id", handlers.DeletePlan) // Delete a plan

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

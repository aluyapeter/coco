package main

import (
	"log"

	"github.com/aluyapeter/coco/config"

	"github.com/aluyapeter/coco/handlers"

	"github.com/aluyapeter/coco/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	//init database conn
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	//Gin router with default middleware (logger and recovery)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
		}

		c.Next()
	})

	//Setup API routes
	setupRoutes(router, db)

	// Step 5: Add health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Task API is running successfully",
			"version": "1.0.0",
		})
	})

	port := "9000"
	log.Printf("🚀 Starting Task API server on port %s...", port)
	log.Printf("📍 Health check: http://localhost:%s/health", port)
	log.Printf("📋 API endpoints: http://localhost:%s/api/v1/tasks", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// setupRoutes configures all API routes
func setupRoutes(router *gin.Engine, db *gorm.DB) {
	// Create repository and handler instances
	taskRepo := repository.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// Create API v1 route group
	api := router.Group("/api/v1")
	{
		// Task routes group
		tasks := api.Group("/tasks")
		{
			// GET /api/v1/tasks/stats - Get task statistics (must be before /:id route)
			tasks.GET("/stats", taskHandler.GetTaskStats)
			
			// POST /api/v1/tasks - Create a new task
			tasks.POST("", taskHandler.CreateTask)
			
			// GET /api/v1/tasks - Get all tasks (supports ?status=completed or ?status=pending)
			tasks.GET("", taskHandler.GetTasks)
			
			// GET /api/v1/tasks/:id - Get a specific task by ID
			tasks.GET("/:id", taskHandler.GetTask)
			
			// PUT /api/v1/tasks/:id - Update an existing task
			tasks.PUT("/:id", taskHandler.UpdateTask)
			
			// DELETE /api/v1/tasks/:id - Delete a task (soft delete)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
}
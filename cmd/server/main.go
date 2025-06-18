package main

import (
	"log"
	"os"

	v1 "github.com/Ik-cyber/markdown-ai-app/api/v1"
	"github.com/Ik-cyber/markdown-ai-app/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	database.Connect()
	defer database.Close()

	// Set up Gin router
	router := gin.Default()

	// API v1 routes
	v1Group := router.Group("/api/v1")
	v1.RegisterRoutes(v1Group)

	// Get port from env or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}

	log.Println("🚀 Server running on port", port)
	router.Run(":" + port)
}

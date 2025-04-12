package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilhamhaikal/Kategori_books.git/config"
	"github.com/ilhamhaikal/Kategori_books.git/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file only in development
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found")
		}
	}

	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

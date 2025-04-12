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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Set Gin mode based on environment
	if os.Getenv("RAILWAY_ENVIRONMENT") != "" {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Running in production mode")
	} else {
		// Load .env in development
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found, using environment variables")
		}
		log.Println("Running in development mode")
	}

	// Print environment variables (excluding sensitive data)
	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))
	log.Printf("DATABASE_URL set: %v", os.Getenv("DATABASE_URL") != "")
	log.Printf("PORT: %s", os.Getenv("PORT"))

	r := gin.Default()

	// Add basic health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Book Category API is running",
		})
	})

	// Connect to database with retries
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err := config.ConnectDB()
		if err == nil {
			log.Println("Database connected successfully")
			break
		}
		if i == maxRetries-1 {
			log.Fatal("Failed to connect to database after", maxRetries, "attempts:", err)
		}
		log.Printf("Database connection attempt %d failed: %v. Retrying...", i+1, err)
	}

	// Run migrations
	if err := config.RunMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrations completed successfully")

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

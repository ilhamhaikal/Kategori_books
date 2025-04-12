package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhamhaikal/Kategori_books.git/config"
	"github.com/ilhamhaikal/Kategori_books.git/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}

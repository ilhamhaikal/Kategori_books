package routes

import (
	"github.com/ilhamhaikal/Kategori_books.git/controllers"
	"github.com/ilhamhaikal/Kategori_books.git/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Public routes
		api.POST("/users/register", controllers.Register)
		api.POST("/users/login", controllers.Login)

		// Protected routes
		auth := api.Group("/")
		auth.Use(middleware.JWTAuth())

		// Book routes
		auth.GET("/books", controllers.GetBooks)
		auth.POST("/books", controllers.CreateBook)
		auth.GET("/books/:id", controllers.GetBookByID)
		auth.PUT("/books/:id", controllers.UpdateBook)
		auth.DELETE("/books/:id", controllers.DeleteBook)

		// Category routes
		auth.GET("/categories", controllers.GetCategories)
		auth.POST("/categories", controllers.CreateCategory)
		auth.GET("/categories/:id", controllers.GetCategoryByID)
		auth.PUT("/categories/:id", controllers.UpdateCategory)
		auth.DELETE("/categories/:id", controllers.DeleteCategory)
		auth.GET("/categories/:id/books", controllers.GetBooksByCategory)
	}
}

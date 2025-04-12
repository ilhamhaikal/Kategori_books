package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilhamhaikal/Kategori_books.git/config"
	"github.com/ilhamhaikal/Kategori_books.git/models"
)

func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate release year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	// Convert thickness based on total_page
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	// Get username from JWT token for audit trail
	username := c.GetString("username")
	now := time.Now()

	err := config.DB.QueryRow(`
        INSERT INTO books (
            title, description, image_url, release_year, 
            price, total_page, thickness, category_id,
            created_at, created_by
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING id`,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, book.Thickness, book.CategoryID,
		now, username).Scan(&book.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	err = config.DB.QueryRow(`
        SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id 
        FROM books WHERE id = $1`, id).
		Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	result, err := config.DB.Exec(`
        UPDATE books 
        SET title=$1, description=$2, image_url=$3, release_year=$4, 
            price=$5, total_page=$6, thickness=$7, category_id=$8,
            modified_at=CURRENT_TIMESTAMP, modified_by=$9
        WHERE id=$10`,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, book.Thickness, book.CategoryID,
		c.GetString("username"), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	book.ID = id
	c.JSON(http.StatusOK, book)
}

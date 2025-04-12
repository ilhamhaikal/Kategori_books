package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	// Try to connect to database
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	return nil
}

func RunMigrations() error {
	// Create users table
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(100) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            created_by VARCHAR(100),
            modified_at TIMESTAMP,
            modified_by VARCHAR(100)
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	// Create categories table
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS categories (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            created_by VARCHAR(100),
            modified_at TIMESTAMP,
            modified_by VARCHAR(100)
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create categories table: %v", err)
	}

	// Create books table
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS books (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            image_url TEXT,
            release_year INTEGER CHECK (release_year BETWEEN 1980 AND 2024),
            price INTEGER,
            total_page INTEGER,
            thickness VARCHAR(50),
            category_id INTEGER REFERENCES categories(id),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            created_by VARCHAR(100),
            modified_at TIMESTAMP,
            modified_by VARCHAR(100)
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create books table: %v", err)
	}

	return nil
}

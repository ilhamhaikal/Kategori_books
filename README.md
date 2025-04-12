# Book Category API Documentation

## Overview
This is a REST API service for managing books and their categories. The API provides endpoints for authentication, book management, and category management with JWT authentication for secure access.

## Authentication
To use protected endpoints, you need to:
1. Login using POST /api/users/login
2. Include the received JWT token in Authorization header

## Available Endpoints

### User Authentication
# Book Category API

REST API for managing books and categories using Go, Gin, and PostgreSQL.

## Features
- JWT Authentication
- Book Management
- Category Management
- PostgreSQL Database

## API Endpoints
### Authentication
- POST /api/users/login - Login user

### Books
- GET /api/books - List all books
- POST /api/books - Create new book
- GET /api/books/:id - Get book details
- DELETE /api/books/:id - Delete book

### Categories
- GET /api/categories - List all categories
- POST /api/categories - Create new category
- GET /api/categories/:id - Get category details
- DELETE /api/categories/:id - Delete category
- GET /api/categories/:id/books - Get books by category

## Setup
1. Clone repository
2. Copy .env.example to .env and configure database
3. Run migrations:
```bash
sql-migrate up
```
4. Run application:
```bash
go run main.go
```

## Deployment
This application is deployed on Railway.
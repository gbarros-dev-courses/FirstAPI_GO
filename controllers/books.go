// controllers/books.go

package controllers

import (
	"net/http"
  "github.com/gin-gonic/gin"
  "github.com/rahmanfadhil/gin-bookstore/models"
)

type CreateBookInput struct {
	Title		string	`json:"title" bingind:"required"`
	Author	string	`json:"author" bingind:"required"`
}

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	var book models.Book

	if err:= models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validadte input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
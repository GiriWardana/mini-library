package controllers

import (
	"net/http"

	"mini-library-backend/models"
	"github.com/gin-gonic/gin"
)

// GET /api/books
func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := bookDB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GET /api/books/:id
func GetBookDetail(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := bookDB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// POST /api/books
func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := bookDB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// PUT /api/books/:id
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bookDB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	book.Stock = input.Stock

	if err := bookDB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DELETE /api/books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := bookDB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := bookDB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

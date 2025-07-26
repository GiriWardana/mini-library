package controllers

import (
	"mini-library-backend/models"
	"net/http"
	"strconv"
	"time"
	"log"

	"github.com/gin-gonic/gin"
)


// POST /api/borrow/:id
func BorrowBook(c *gin.Context) {
	userID := c.GetUint("user_id") // Ambil dari JWT middleware
	bookID, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	if err := borrowDB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.Stock < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book out of stock"})
		return
	}

	// Create borrow record
	borrow := models.Borrow{
		UserID:   userID,
		BookID:   uint(bookID),
		Returned: false,
	}
	if err := borrowDB.Create(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book"})
		return
	}

	// Reduce stock
	book.Stock--
	borrowDB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed"})
}

// POST /api/borrow/return/:id
func ReturnBook(c *gin.Context) {
	userID := c.GetUint("user_id")
	log.Println("🚀 ~ userID:", userID)
	borrowID := c.Param("id")

	var borrow models.Borrow
	if err := borrowDB.First(&borrow, borrowID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrow record not found"})
		return
	}

	if borrow.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if borrow.Returned {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book already returned"})
		return
	}

	// Mark as returned
	now := time.Now()
	borrow.Returned = true
	borrow.ReturnedAt = &now
	borrowDB.Save(&borrow)

	// Increase book stock
	var book models.Book
	borrowDB.First(&book, borrow.BookID)
	book.Stock++
	borrowDB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Book returned"})
}

// GET /api/borrow/history (user)
func GetUserHistory(c *gin.Context) {
	userID := c.GetUint("user_id")
	var history []models.Borrow
	borrowDB.Preload("Book").Where("user_id = ?", userID).Find(&history)
	c.JSON(http.StatusOK, history)
}

// GET /api/admin/history (admin only)
func GetAllHistory(c *gin.Context) {
	var history []models.Borrow
	borrowDB.Preload("User").Preload("Book").Find(&history)
	c.JSON(http.StatusOK, history)
}

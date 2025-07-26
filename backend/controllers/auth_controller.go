package controllers

import (
	"net/http"
	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"mini-library-backend/models"
)

// var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // secret key from env

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]struct {
	Password string
	Role     string
}{
	"admin": {"password", "admin"},
	"user1": {"password", "user"},
}

func Register(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var existing models.User
	if err := userDB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	newUser := models.User{
		Username: req.Username,
		Password: req.Password,
		Role:     "user",
	}
	if err := userDB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Register success", "role": newUser.Role})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var user models.User
	if err := userDB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.Username,
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"iat":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("mini-library-secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"role":  user.Role,
	})
}


func SeedDefaultUsers() {
	seed := []models.User{
		{Username: "admin", Password: "password", Role: "admin"},
		{Username: "user1", Password: "password", Role: "user"},
	}

	for _, u := range seed {
		var existing models.User
		if err := userDB.Where("username = ?", u.Username).First(&existing).Error; err != nil {
			userDB.Create(&u)
		}
	}
}

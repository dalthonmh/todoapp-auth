package controllers

import (
	"net/http"

	"github.com/dalthonmh/todoapp-auth/models"
	"github.com/dalthonmh/todoapp-auth/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Hash password
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing password"})
			return
		}
		input.Password = string(hashed)

		// Create user
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		var user models.User
		if err := db.First(&user, "username = ?", input.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}

		// Generate token
		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

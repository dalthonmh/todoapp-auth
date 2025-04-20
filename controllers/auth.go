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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Hashear contraseña
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la contraseña"})
			return
		}
		input.Password = string(hashed)

		// Crear usuario
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Usuario ya existe"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado con éxito"})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		var user models.User
		if err := db.First(&user, "username = ?", input.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}

		// Generar token
		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

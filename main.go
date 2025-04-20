package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dalthonmh/todoapp-auth/models"
	"github.com/dalthonmh/todoapp-auth/routes"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Migración automática
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.AuthRoutes(r, db)

	r.Run(":8080")
}

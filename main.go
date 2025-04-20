package main

import (
	"log"

	"github.com/dalthonmh/todoapp-auth/models"
	"github.com/dalthonmh/todoapp-auth/routes"
	"github.com/dalthonmh/todoapp-auth/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db := utils.ConnectWithRetry()

	// Migración automática
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error al migrar el modelo:", err)
	}

	// Configuración del servidor Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.AuthRoutes(r, db)

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

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

	// Automatic migration
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error migrating the model:", err)
	}

	// Gin server configuration
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	// Check status
	r.GET("/api/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth microservice running",
		})
	})

	routes.AuthRoutes(r, db)

	// Init server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

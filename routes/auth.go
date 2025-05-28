package routes

import (
	"github.com/dalthonmh/todoapp-auth/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/api/auth/register", controllers.Register(db))
	r.POST("/api/auth/login", controllers.Login(db))
}

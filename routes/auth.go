package routes

import (
	"github.com/dalthonmh/todoapp-auth/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.Register(db))
	r.POST("/login", controllers.Login(db))
}

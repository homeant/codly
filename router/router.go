package router

import (
	"codly/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	var router = gin.Default()
	router.GET("/app", handlers.AppHandler(db))
	router.PUT("/admin-user", handlers.CreateAdminUserHandler(db))
	return router
}

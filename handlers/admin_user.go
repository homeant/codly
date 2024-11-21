package handlers

import (
	"codly/model"
	"codly/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAdminUserHandler(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		a := service.NewAdminUserService(db)
		a.CreateAdminUser(model.AdminUser{Username: "hello", Password: "password"})
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

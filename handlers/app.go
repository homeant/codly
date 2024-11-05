package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppHandler(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

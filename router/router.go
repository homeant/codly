package router

import (
	"codly/api"
	"codly/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)

	authenticated := r.Group("/api")
	authenticated.Use(middleware.JWTMiddleware())

	authenticated.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

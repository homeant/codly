package router

import (
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/api"
	"github.com/homeanter/codly/middleware"
)

func New() *gin.Engine {
	r := gin.Default()
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)

	authenticated := r.Group("/api")
	authenticated.Use(middleware.JWTMiddleware())

	authenticated.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

package main

import (
	"codly/handlers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter() {
	router = gin.Default()
	router.GET("/app", handlers.AppHandler)
}

func GetRouter() *gin.Engine {
	return router
}

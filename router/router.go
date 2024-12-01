package router

import (
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/api"
)

func Init() *gin.Engine {
	r := gin.New()
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	return r
}

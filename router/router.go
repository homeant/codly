package router

import (
	"codly/api"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.GET("/app", api.Login)
	return r
}

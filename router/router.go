package router

import (
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/api"
)

func Init(r *gin.Engine) {
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
}

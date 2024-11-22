package api

import (
	"codly/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	result, _ := service.AdminUserService.GetAdminUser("hello")
	c.JSON(200, result)
}

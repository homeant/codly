package api

import (
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/model"
	"github.com/homeanter/codly/service"
)

func Register(c *gin.Context) {
	data := model.AdminUserRegister{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := service.AdminUserService.Register(&data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, model.StatusMessage{
		Message: "success",
	})
	return
}

func Login(c *gin.Context) {
	loginData := model.AdminUserLogin{}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := service.AdminUserService.Login(&loginData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	if result == nil {
		c.JSON(404, gin.H{"error": "not found"})
	}
	c.JSON(200, result)
	return
}

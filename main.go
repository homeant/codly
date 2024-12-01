package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/config"
	"github.com/homeanter/codly/router"
	"github.com/homeanter/codly/utils"
	"os"
)

func main() {
	var r = router.Init()
	logFilePath := "./logs/codly.log"
	err := utils.CreateDir(logFilePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	gin.DefaultWriter = logFile
	fmt.Println("Codly is running on port: ", config.App.Port)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

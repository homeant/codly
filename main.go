package main

import (
	"codly/config"
	"codly/logging"
	"codly/middleware"
	"codly/router"
	"codly/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	router.Init(r)

	logFilePath := "./logs/codly.log"
	err := utils.CreateDir(logFilePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(logFile)
	logging.DefaultLogger.SetOutput(io.MultiWriter(logFile, os.Stdout))
	logging.DefaultLogger.SetLevel(logging.Debug)
	logging.DefaultLogger.Infof("%s is running on port: %d", config.Config.App.Name, config.App.Port)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

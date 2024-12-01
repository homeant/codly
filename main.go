package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/config"
	"github.com/homeanter/codly/router"
	"github.com/homeanter/codly/utils"
	"io"
	"os"
)

func main() {
	r := gin.Default()
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
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(logFile)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format("2006-01-02 15:04:05.000"),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
	}))
	fmt.Println("Codly is running on port: ", config.App.Port)
	router.Init(r)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

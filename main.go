package main

import (
	"codly/config"
	"codly/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	var r = router.Init()
	logFile, err := os.Create("./logs/codly.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = logFile
	fmt.Println("Codly is running on port: ", config.App.Port)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

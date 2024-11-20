package main

import (
	"codly/datastore"
	"codly/model"
	"codly/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	config := model.Config{}
	fileBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(fileBytes, &config)
	if err != nil {
		return
	}
	db, err := datastore.InitDB(config.Database)
	if err != nil {
		return
	}
	var r = router.InitRouter(db)
	logFile, err := os.Create("./logs/codly.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = logFile
	r.Use(gin.Logger())
	fmt.Println("Codly is running on port: ", config.App.Port)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

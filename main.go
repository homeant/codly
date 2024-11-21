package main

import (
	"codly/datastore"
	"codly/model"
	"codly/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
	"time"
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
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)

	//// SetMaxOpenConns 设置数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接的最大重用时间
	sqlDB.SetConnMaxLifetime(time.Hour)

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

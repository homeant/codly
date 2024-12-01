package main

import (
	"fmt"
	"github.com/homeanter/codly/config"
	"github.com/homeanter/codly/middleware"
	"github.com/homeanter/codly/router"
	"log"
)

func main() {
	r := router.New()
	// 应用日志中间件到所有路由
	r.Use(middleware.LoggerMiddleware())
	log.Println("Codly is running on port: ", config.App.Port)
	_ = r.Run(fmt.Sprintf(":%d", config.App.Port))
}

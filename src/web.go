package main

import (
	"blog/router"
	"blog/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	app := gin.New()
	app.Use(gin.Recovery())

	baseGroup := app.Group("/api")
	{
		router.InitRoutes(baseGroup)
	}

	return app
}

func RunServer() {
	config := utils.Config()

	app := InitApp()

	log.Println("Starting server...")

	app.Run(config.ServerDsn)
}

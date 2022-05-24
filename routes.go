package main

import (
	"go-web-boilerplate/controllers"
	"go-web-boilerplate/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/hello", controllers.Hello)
		}
	}

	router.Use(static.Serve("/", static.LocalFile("./ui/dist", false)))
	return router
}

package main

import (
	"fmt"
	"go-mongodb-api/controllers"
	"go-mongodb-api/middleware"
	"go-mongodb-api/services"
	"go-mongodb-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("[API] Started on", utils.Config().Server.Host+":"+utils.Config().Server.Port)

	services.Database()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/token", controllers.Token)

	router.GET("/articles", middleware.AuthUser(), controllers.GetArticles)
	router.POST("/articles", middleware.AuthUser(), controllers.PostArticles)

	router.Run(":" + utils.Config().Server.Port)
}

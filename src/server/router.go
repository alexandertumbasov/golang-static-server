package server

import (
	"github.com/gin-gonic/gin"
	"golang-static-server/src/controllers"
	"golang-static-server/src/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		fileController := new(controllers.File)

		v1.GET("/file", fileController.SaveFileHandler)

	}

	return router
}

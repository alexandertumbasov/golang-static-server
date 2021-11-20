package server

import (
	"github.com/gin-gonic/gin"
	"golang-static-server/middlewares"
	"golang-static-server/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		userController := new(controllers.UserController)

		v1.GET("/", userController.GetAll)

	}

	return router
}

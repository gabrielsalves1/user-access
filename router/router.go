package router

import (
	"user-access/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:userId", controller.GetUser)

	router.POST("/systems", controller.CreateSystem)
	router.GET("/systems", controller.GetAllSystems)
	router.GET("/systems/:systemId", controller.GetSystem)

	router.POST("/user-systems", controller.AddUserToSystem)

	router.Run("localhost:8080")
}

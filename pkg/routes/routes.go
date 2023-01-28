package routes

import (
	"user-access/pkg/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	routes := gin.Default()

	routes.POST("/users", controller.CreateUser)
	routes.GET("/users", controller.GetAllUsers)
	routes.GET("/users/:userId", controller.GetUser)

	routes.POST("/systems", controller.CreateSystem)
	routes.GET("/systems", controller.GetAllSystems)
	routes.GET("/systems/:systemId", controller.GetSystem)

	routes.POST("/user-systems", controller.AddUserToSystem)

	routes.Run("localhost:8080")
}

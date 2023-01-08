package router

import (
	userController "user-access/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/users", userController.GetAllUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:userId", userController.GetUser)

	router.Run("localhost:8080")
}

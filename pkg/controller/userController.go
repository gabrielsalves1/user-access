package controller

import (
	"log"
	"net/http"
	"user-access/pkg/database"
	"user-access/pkg/models"

	"github.com/gin-gonic/gin"
)

type SetUserToSystem struct {
	UserId   int `json:"user_id"`
	SystemId int `json:"system_id"`
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.Create(&user); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Preload("Systems").Find(&users); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	var user models.User

	if err := database.DB.Preload("Systems").First(&user, userId); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusOK, &user)
}

func AddUserToSystem(c *gin.Context) {
	body := SetUserToSystem{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("Search user id: %c\n, system_id: %c", body.UserId, body.SystemId)

	var user models.User

	if err := database.DB.Preload("Systems").Find(&user, body.UserId); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	var system models.System

	if err := database.DB.Find(&system, body.SystemId); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	database.DB.Model(&user).Association("Systems").Append(&system)
	c.JSON(http.StatusOK, &user)
}

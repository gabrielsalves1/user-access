package controller

import (
	"net/http"
	"user-access/config"
	"user-access/models"

	"github.com/gin-gonic/gin"
)

type AddUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Number  int    `json:"number"`
	Company string `json:"company"`
	Team    string `json:"team"`
}

func CreateUser(c *gin.Context) {
	body := AddUser{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	user.Name = body.Name
	user.Email = body.Email
	user.Address = body.Address
	user.Number = body.Number
	user.City = body.City
	user.State = body.State
	user.Country = body.Country
	user.Company = body.Company
	user.Team = body.Team

	if result := config.Database().Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if result := config.Database().Preload("Systems").Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	var user models.User

	if result := config.Database().Preload("Systems").First(&user, userId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &user)
}

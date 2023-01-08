package userController

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
	Number  int    `json:"number"`
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if result := config.Database().Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, users)
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

	if result := config.Database().Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	var user models.User

	if result := config.Database().First(&user, userId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &user)
}

package controller

import (
	"net/http"
	"user-access/config"
	"user-access/models"

	"github.com/gin-gonic/gin"
)

type AddSystem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func CreateSystem(c *gin.Context) {
	body := AddSystem{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var system models.System
	system.Name = body.Name
	system.Url = body.Url

	if result := config.Database().Create(&system); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &system)
}

func GetAllSystems(c *gin.Context) {
	var systems []models.System

	if result := config.Database().Find(&systems); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, systems)
}

func GetSystem(c *gin.Context) {
	systemId := c.Param("systemId")
	var system models.System

	if result := config.Database().First(&system, systemId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, system)
}

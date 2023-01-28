package controller

import (
	"net/http"
	"user-access/pkg/database"
	"user-access/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateSystem(c *gin.Context) {
	var system models.System

	if err := c.ShouldBindJSON(&system); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := database.DB.Create(&system); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusCreated, system)
}

func GetAllSystems(c *gin.Context) {
	var systems []models.System

	if err := database.DB.Find(&systems); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusOK, systems)
}

func GetSystem(c *gin.Context) {
	systemId := c.Param("systemId")
	var system models.System

	if err := database.DB.First(&system, systemId); err.Error != nil {
		c.AbortWithError(http.StatusNotFound, err.Error)
		return
	}

	c.JSON(http.StatusOK, system)
}

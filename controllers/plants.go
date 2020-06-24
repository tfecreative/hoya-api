package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tfecreative.com/hoya-api/models"
)

func FindPlants(c *gin.Context) {
	var plants []models.Plant
	models.DB.Find(&plants)

	c.JSON(http.StatusOK, gin.H{"data": plants})
}
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfecreative/hoya-api/api/models"
)

type CreatePlantInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePlantInput struct {
	Name string `json:"name"`
}

func CreatePlant(c *gin.Context) {
	var input CreatePlantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plant := models.Plant{Name: input.Name}
	models.DB.Create(&plant)

	c.JSON(http.StatusOK, gin.H{"data": plant})
}

func DeletePlant(c *gin.Context) {
	var plant models.Plant
	if err := models.DB.Where("id = ?", c.Param("id")).First(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plant not found!"})
		return
	}

	models.DB.Delete(&plant)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindPlant(c *gin.Context) {
	var plant models.Plant

	if err := models.DB.Where("id = ?", c.Param("id")).First(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plant not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plant})
}

func FindPlants(c *gin.Context) {
	var plants []models.Plant
	models.DB.Find(&plants)

	c.JSON(http.StatusOK, gin.H{"data": plants})
}

func UpdatePlant(c *gin.Context) {
	var plant models.Plant
	if err := models.DB.Where("id = ?", c.Param("id")).First(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plant not found!"})
		return
	}

	var input UpdatePlantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&plant).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": plant})
}

package main


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tfecreative/hoya-api/controllers"
	"github.com/tfecreative/hoya-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "OK"})
	})

	r.GET("/plants", controllers.FindPlants)

	r.POST("/plants", controllers.CreatePlant)

	r.GET("/plants/:id", controllers.FindPlant)

	r.PATCH("/plants/:id", controllers.UpdatePlant)

	r.DELETE("/plants/:id", controllers.DeletePlant)

	r.Run(":8000")
}
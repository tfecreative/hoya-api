package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tfecreative/hoya-api/controllers"
	"github.com/tfecreative/hoya-api/models"

	"github.com/gin-contrib/cors"
	c "github.com/tfecreative/hoya-api/config"
)

func main() {
	c.LoadConfig()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

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

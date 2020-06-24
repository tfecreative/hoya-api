package main


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tfecreative.com/hoya-api/controllers"
	"tfecreative.com/hoya-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "OK"})
	})

	r.GET("/plants", controllers.FindPlants)

	r.Run()
}
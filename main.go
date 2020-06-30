package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/tfecreative/hoya-api/api/controllers"
	"github.com/tfecreative/hoya-api/api/models"

	c "github.com/tfecreative/hoya-api/config"
)

func main() {
	c.LoadConfig()

	r := mux.NewRouter()

	models.ConnectDataBase()

	r.HandleFunc("/plants", controllers.PlantsHandler).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	srv.ListenAndServe()
}

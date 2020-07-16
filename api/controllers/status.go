package controllers

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/tfecreative/hoya-api/api/models"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	client := models.GetDb()
	collection := client.Collection("plants")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	plant := &models.Plant{Name: "Hoya carnosa"}
	_, err := collection.InsertOne(ctx, plant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// A very simple health check
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

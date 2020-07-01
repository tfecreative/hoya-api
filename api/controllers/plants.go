package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/tfecreative/hoya-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func PlantsHandler(w http.ResponseWriter, r *http.Request) {
	client := models.GetDb()
	collection := client.Collection("plants")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var plants []models.Plant
	cur, err := collection.Find(ctx, bson.M{}) // k
	if err != nil {
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var plant models.Plant
		cur.Decode(&plant)
		plants = append(plants, plant)
	}
	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(w).Encode(plants)
}

func CreatePlantHandler(w http.ResponseWriter, r *http.Request) {
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

	var createdPlant models.Plant
	err = collection.FindOne(ctx, bson.M{"name": "Hoya carnosa"}).Decode(&createdPlant)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPlant)
}

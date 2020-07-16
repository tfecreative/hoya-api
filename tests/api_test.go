package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tfecreative/hoya-api/api/controllers"
	"github.com/tfecreative/hoya-api/api/models"
)

func TestPlants(t *testing.T) {
	req, err := http.NewRequest("GET", "/plants", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.PlantsHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)

	var data []models.Plant
	json.NewDecoder(rr.Body).Decode(&data)
	assert.GreaterOrEqual(t, len(data), 0)
}

func TestCreatePlant(t *testing.T) {
	req, err := http.NewRequest("POST", "/plants", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreatePlantHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusCreated)
}

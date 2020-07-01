package tests

import (
	"context"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/tfecreative/hoya-api/api/models"
	c "github.com/tfecreative/hoya-api/config"

	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestDbIsTestDb(t *testing.T) {
	dbStr := viper.Get("TEST_MONGO_DB_CONNECTION_URI")
	assert.Equal(t, "mongodb://mongo:27017/", dbStr)
}

func dropDatabase() {
	db := models.GetDb()
	collection := db.Collection("plants")
	err := collection.Drop(context.Background())
	if err != nil {
		log.Warn("Failed to drop database!")
	}
}

func setup() {
	c.LoadConfig()
	dropDatabase()
}

func shutdown() {
	dropDatabase()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

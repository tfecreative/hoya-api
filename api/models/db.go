package models

import (
	"context"
	"flag"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	testing := flag.Lookup("test.v") != nil
	var dbUri = ""
	if testing {
		dbUri = viper.GetString("TEST_MONGO_DB_CONNECTION_URI")
	} else {
		dbUri = viper.GetString("MONGO_DB_CONNECTION_URI")
	}

	// Create db client
	clientOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDb() *mongo.Database {
	client := GetClient()
	db := client.Database("hoya")
	return db
}

package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetDefault("MONGO_DB_CONNECTION_URI", "mongodb://mongo:27017/")
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning("Failed to load .env file")
	}
}

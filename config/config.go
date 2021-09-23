package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig init viper config
func InitConfig() {
	// config file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	// config file path
	viper.AddConfigPath("./config")
	// try to load config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fail to read config file: ", err)
	}
}

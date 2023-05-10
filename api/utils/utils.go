package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("MIGRATE", true)
	viper.SetDefault("ENVIRONMENT", "development")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicln(fmt.Errorf("CONFIG FILE NOT FOUND"))
		} else {
			log.Panicln(fmt.Errorf("Fatal error config file: %s", err))
		}
	}
}

func GetPort() string {
	return strconv.Itoa(viper.GetInt("PORT"))
}

package clio

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("clio")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
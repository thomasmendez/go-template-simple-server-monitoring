package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"GO_SERVER_ADDRESS"`
	Environment   string `mapstructure:"GO_ENV"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.SetDefault("GO_SERVER_ADDRESS", "0.0.0.0:8081")
	viper.SetDefault("GO_ENV", "development")

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error while reading config file: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode config file to struct, err: %v", err)
		return
	}

	return
}

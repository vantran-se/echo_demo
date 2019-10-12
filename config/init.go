package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Schema struct {
	Mongo struct {
		Host     string `mapstructure:"Host"`
		Username string `mapstructure:"Username"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`

	Encryption struct {
		JWTSecret string `mapstructure:"Jwt"`
	} `mapstructure:"Encrypt"`
}

var Config Schema

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("Fatal error: %s \n", err))
	}
}

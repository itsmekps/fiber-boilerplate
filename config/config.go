package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `envconfig:"PORT" default:"3000"`
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return v, nil
}

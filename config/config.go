package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return v, nil
}

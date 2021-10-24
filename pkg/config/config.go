package config

import (
	"time"

	viper "github.com/spf13/viper"
)

type Configer interface {
	Get(string) interface{}
	GetString(string) string
	GetDuration(string) time.Duration
}

func NewConfig() Configer {
	viper.SetConfigFile(".config.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetViper()
}

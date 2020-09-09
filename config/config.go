package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var C = Conf()

func Conf() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config/")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return v
}

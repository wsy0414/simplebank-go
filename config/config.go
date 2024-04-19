package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Driver string `mapstructure:"driverName"`
	Source string `mapstructure:"source"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Database Database `mapstructure:"db"`
	Server   Server   `mapstructure:"server"`
}

var ConfigVal Config

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file error")
	}

	err = viper.Unmarshal(&ConfigVal)
	if err != nil {
		panic("unmarshal config failed")
	}
}

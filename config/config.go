package config

import (
	"fmt"
	"simplebank/flags"

	"github.com/spf13/viper"
)

type Database struct {
	Driver string `mapstructure:"driverName"`
	Source string `mapstructure:"source"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Grpc struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Database Database `mapstructure:"db"`
	Server   Server   `mapstructure:"server"`
	Grpc     Grpc     `mapstructure:"grpc"`
}

var ConfigVal Config

func LoadConfig(prePath string) {

	viper.AddConfigPath(prePath)
	configName := "app"
	if flags.ENV != "" {
		configName += ("_" + flags.ENV)
	}
	fmt.Println("configName:", configName)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ConfigVal)
	if err != nil {
		panic("unmarshal config failed")
	}
}

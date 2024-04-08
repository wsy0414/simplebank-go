package config

import "github.com/spf13/viper"

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

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}

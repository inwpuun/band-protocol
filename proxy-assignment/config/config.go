package config

import (
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		ServerConfig        ServerConfig        `mapstructure:"server_config"`
		Server2ServerConfig Server2ServerConfig `mapstructure:"server2server"`
	}

	ServerConfig struct {
		Port int `mapstructure:"port"`
	}

	Server2ServerConfig struct {
		Url string `mapstructure:"url"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}

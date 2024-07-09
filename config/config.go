package config

import (
	"github.com/spf13/viper"

	"strings"
	"sync"
)

type (
	Config struct {
		Server   *Server
		Database *Database
	}

	Server struct {
		Port int
	}

	Database struct {
		User     string
		Password string
		Protocol string
		Path     string
		DBName   string
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
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}

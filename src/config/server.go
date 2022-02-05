package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Server struct {
		Port uint16 `mapstructure:"port"`
	}
}

var serverConfig *ServerConfig

func initServerConfig() {
	serverConfig = &ServerConfig{}
	err := viper.Unmarshal(serverConfig)
	if err != nil {
		log.Printf("Cannot unmarshal config to server config: %s", err)
	}
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}

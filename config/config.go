package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `json:",squash"`
}

type Server struct {
	Host string `json:"server_host"`
	Port string `json:"server_port"`
}

func SetDefaults() {
	viper.SetDefault("SERVER_HOST", "localhost")
	viper.SetDefault("SERVER_PORT", 8000)
}

func NewConfigFromFile() (*Config, error) {
	config := &Config{}

	SetDefaults()
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

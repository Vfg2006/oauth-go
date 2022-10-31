package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:",squash"`
}

type Server struct {
	Host string `mapstructure:"server_host"`
	Port string `mapstructure:"server_port"`
}

func SetDefaults() {
	viper.SetDefault("server_host", "localhost")
	viper.SetDefault("server_port", 8000)
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

	err = viper.Unmarshal(&config, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	))
	if err != nil {
		return nil, err
	}

	return config, nil
}

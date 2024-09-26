package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct{}
	Secret struct {
		ChatGPTToken string `mapstructure:"CHAT_GPT_TOKEN"`
	}
)

func LoadConfig(configPath, secretPath string) (*Config, *Secret, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	viper.Reset()

	// Load Secret
	viper.SetConfigFile(secretPath)
	var secret Secret
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading secret file: %w", err)
	}
	if err := viper.Unmarshal(&secret); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, &secret, nil
}

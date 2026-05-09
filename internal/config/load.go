package config

import (
	"github.com/spf13/viper"
)

// LoadConfig unmarshals viper's loaded config into a struct
func LoadConfig() (*Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

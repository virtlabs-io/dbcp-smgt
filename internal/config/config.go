package config

import (
	"github.com/spf13/viper"
)

// GlobalConfig defines the global configuration fields.
type GlobalConfig struct {
	LogLevel    string         `mapstructure:"log_level"`
	Environment string         `mapstructure:"environment"`
	Modules     []ModuleConfig `mapstructure:"modules"`
}

// ModuleConfig defines per‑module configuration.
type ModuleConfig struct {
	Name       string `mapstructure:"name"`
	Enabled    bool   `mapstructure:"enabled"`
	Path       string `mapstructure:"path"`
	ConfigFile string `mapstructure:"config_file"`
}

// Config aggregates all configuration sections.
type Config struct {
	Global GlobalConfig `mapstructure:"global"`
	// You can add more sections if needed.
}

// LoadConfig loads and unmarshals configuration from a file.
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

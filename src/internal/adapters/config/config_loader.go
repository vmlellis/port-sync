package config

import (
	"github.com/BurntSushi/toml"
)

// Config represents the configuration structure
type Config struct {
	PortsFile     string `toml:"ports_file"`
	LoadOnStartup bool   `toml:"load_on_startup"`
}

// LoadConfig reads configuration from a TOML file
func LoadConfig(configPath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

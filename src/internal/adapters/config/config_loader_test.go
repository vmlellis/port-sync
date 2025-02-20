package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmlellis/port-sync/src/internal/adapters/config"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	configData := `ports_file = "data/ports.json"
load_on_startup = true
`

	tmpFile, err := os.CreateTemp("", "config.toml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(configData))
	assert.NoError(t, err)
	tmpFile.Close()

	// Load the configuration
	config, err := config.LoadConfig(tmpFile.Name())
	assert.NoError(t, err)

	// Validate the configuration values
	assert.Equal(t, "data/ports.json", config.PortsFile)
	assert.True(t, config.LoadOnStartup)
}

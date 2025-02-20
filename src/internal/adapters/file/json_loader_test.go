package file_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmlellis/port-sync/src/internal/adapters/file"
	"github.com/vmlellis/port-sync/src/internal/adapters/processor"
	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/service"
)

func TestProcessJSONFile(t *testing.T) {
	// Create a temporary JSON file
	jsonData := `{
			"PORT1": {
					"name": "Test Port",
					"city": "Test City",
					"country": "Test Country",
					"coordinates": [12.34, 56.78],
					"province": "Test Province",
					"timezone": "UTC+1",
					"unlocs": ["UNLOC1"],
					"code": "1234"
			}
	}`

	tmpFile, err := os.CreateTemp("", "ports.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(jsonData))
	assert.NoError(t, err)
	tmpFile.Close()

	// Initialize storage and service
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processorService := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{})

	// Process the JSON file
	err = file.ProcessJSONFile(tmpFile.Name(), portService, processorService)
	assert.NoError(t, err)

	// Verify that the port was stored
	retrievedPort, found := portService.GetPort("PORT1")
	assert.True(t, found, "Port should be found in storage")
	assert.Equal(t, "Test Port", retrievedPort.Name)
	assert.Equal(t, "Test City", retrievedPort.City)
	assert.Equal(t, "Test Country", retrievedPort.Country)
	assert.Equal(t, "Test Province", retrievedPort.Province)
	assert.Equal(t, "UTC+1", retrievedPort.Timezone)
	assert.Equal(t, "1234", retrievedPort.Code)
}

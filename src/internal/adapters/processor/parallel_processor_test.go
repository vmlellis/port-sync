package processor_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmlellis/port-sync/src/internal/adapters/processor"
	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/service"
)

func TestParallelProcessor_Process(t *testing.T) {
	// Simulate a JSON payload
	jsonData := `{
			"PORT1": {
					"name": "Test Port 1",
					"city": "Test City 1",
					"country": "Test Country 1",
					"coordinates": [12.34, 56.78],
					"province": "Test Province 1",
					"timezone": "UTC+1",
					"unlocs": ["UNLOC1"],
					"code": "1234"
			},
			"PORT2": {
					"name": "Test Port 2",
					"city": "Test City 2",
					"country": "Test Country 2",
					"coordinates": [98.76, 54.32],
					"province": "Test Province 2",
					"timezone": "UTC+2",
					"unlocs": ["UNLOC2"],
					"code": "5678"
			}
	}`

	// Create an in-memory reader for JSON data
	reader := bytes.NewReader([]byte(jsonData))

	// Initialize storage and service
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processor := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{BufferedChannelSize: 100, WorkersPoolSize: 5})

	// Process the JSON data
	err := processor.Process(reader)
	assert.NoError(t, err)

	// Verify that the ports were stored
	retrievedPort1, found1 := portService.GetPort("PORT1")
	retrievedPort2, found2 := portService.GetPort("PORT2")

	assert.True(t, found1, "PORT1 should be found in storage")
	assert.Equal(t, "Test Port 1", retrievedPort1.Name)

	assert.True(t, found2, "PORT2 should be found in storage")
	assert.Equal(t, "Test Port 2", retrievedPort2.Name)
}

func TestParallelProcessor_LargeDataset(t *testing.T) {
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processor := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{BufferedChannelSize: 100, WorkersPoolSize: 10})

	// Generate JSON for 1 million ports
	var jsonData bytes.Buffer
	jsonData.WriteString("{")
	for i := 0; i < 1_000_000; i++ {
		if i > 0 {
			jsonData.WriteString(",")
		}
		jsonData.WriteString(fmt.Sprintf(`"PORT%d": {"name": "Test Port %d", "city": "Test City %d", "country": "Test Country %d"}`, i, i, i, i))

	}
	jsonData.WriteString("}")

	reader := bytes.NewReader(jsonData.Bytes())
	err := processor.Process(reader)
	assert.NoError(t, err)

	// Verify a random port is stored
	for _, portID := range []string{"PORT1000", "PORT500000", "PORT999999"} {
		retrievedPort, found := portService.GetPort(portID)
		assert.True(t, found, fmt.Sprintf("%s should be found in storage", portID))
		fmt.Printf("%v", portID)
		fmt.Printf("%v", retrievedPort)
		assert.Equal(t, fmt.Sprintf("Test Port %s", portID[4:]), retrievedPort.Name)
	}
}

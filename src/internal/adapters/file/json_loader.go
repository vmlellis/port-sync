package file

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// ProcessJSONFile reads and processes a JSON file containing port data.
func ProcessJSONFile(filePath string, service contract.PortService) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening bracket
	if _, err := decoder.Token(); err != nil {
		return err
	}

	for decoder.More() {
		var port entity.Port
		key, err := decoder.Token()
		if err != nil {
			return err
		}

		if id, ok := key.(string); ok {
			port.ID = id
		} else {
			return fmt.Errorf("unexpected key type: %T", key)
		}

		if err := decoder.Decode(&port); err != nil {
			return err
		}

		service.SavePort(&port)
	}

	return nil
}

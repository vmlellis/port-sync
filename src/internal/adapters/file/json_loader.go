package file

import (
	"os"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
)

// ProcessJSONFile reads and processes a JSON file containing port data.
func ProcessJSONFile(filePath string, service contract.PortService, processor contract.ProcessorService) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return processor.Process(file)
}

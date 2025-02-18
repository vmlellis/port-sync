package contract

import "github.com/vmlellis/port-sync/src/internal/domain/entity"

// PortRepository defines the contract for persisting port data.
type PortRepository interface {
	// Save inserts or updates a port record in the storage.
	Save(port *entity.Port)

	// Get retrieves a port by its unique identifier.
	Get(id string) (*entity.Port, bool)
}

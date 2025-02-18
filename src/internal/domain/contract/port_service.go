package contract

import "github.com/vmlellis/port-sync/src/internal/domain/entity"

// PortService defines the contract for handling port data.
type PortService interface {
	// SavePort inserts or updates a port record in the storage.
	SavePort(port *entity.Port)

	// GetPort retrieves a port by its unique identifier.
	GetPort(id string) (*entity.Port, bool)
}

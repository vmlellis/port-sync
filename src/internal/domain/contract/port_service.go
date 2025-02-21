package contract

import (
	"context"

	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// PortService defines the contract for handling port data.
type PortService interface {
	// SavePort inserts or updates a port record in the storage.
	SavePort(ctx context.Context, port *entity.Port)

	// GetPort retrieves a port by its unique identifier.
	GetPort(ctx context.Context, id string) (*entity.Port, bool)
}

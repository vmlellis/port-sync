package contract

import (
	"context"

	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// PortRepository defines the contract for persisting port data.
type PortRepository interface {
	// Save inserts or updates a port record in the storage.
	Save(ctx context.Context, port *entity.Port) error

	// Get retrieves a port by its unique identifier.
	Get(ctx context.Context, id string) (*entity.Port, bool)

	// Ping checks the connection
	Ping(ctx context.Context) error

	// Close closes the connection
	Close() error
}

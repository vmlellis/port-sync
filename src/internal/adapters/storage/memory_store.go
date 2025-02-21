package storage

import (
	"context"
	"sync"

	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// MemoryStore implements the PortRepository interface using in-memory storage.
type MemoryStore struct {
	data sync.Map
}

// NewMemoryStore initializes a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// Save inserts or updates a port record in memory.
func (s *MemoryStore) Save(_ context.Context, port *entity.Port) error {
	s.data.Store(port.ID, port)
	return nil
}

// Get retrieves a port by its unique identifier.
func (s *MemoryStore) Get(_ context.Context, id string) (*entity.Port, bool) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, false
	}
	return val.(*entity.Port), true
}

// Ping does nothing for MemoryStore.
func (s *MemoryStore) Ping(_ context.Context) error {
	return nil
}

// Close does nothing for MemoryStore.
func (s *MemoryStore) Close() error {
	return nil
}

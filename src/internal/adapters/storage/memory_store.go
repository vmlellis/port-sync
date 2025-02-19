package storage

import (
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
func (s *MemoryStore) Save(port *entity.Port) {
	s.data.Store(port.ID, port)
}

// Get retrieves a port by its unique identifier.
func (s *MemoryStore) Get(id string) (*entity.Port, bool) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, false
	}
	return val.(*entity.Port), true
}

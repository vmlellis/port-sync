package storage

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"

	"github.com/klauspost/compress/zstd"
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
	var buf bytes.Buffer
	jsonData, _ := json.Marshal(port)

	encoder, _ := zstd.NewWriter(&buf) // Open a writer for each compression
	_, _ = encoder.Write(jsonData)
	encoder.Close()

	s.data.Store(port.ID, buf.Bytes())
}

// Get retrieves a port by its unique identifier.
func (s *MemoryStore) Get(id string) (*entity.Port, bool) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, false
	}

	decoder, _ := zstd.NewReader(bytes.NewReader(val.([]byte)))
	decompressed, _ := io.ReadAll(decoder)
	decoder.Close()

	var port entity.Port
	json.Unmarshal(decompressed, &port)
	return &port, true
}

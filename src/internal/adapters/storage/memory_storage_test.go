package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

func TestMemoryStore_SaveAndGet(t *testing.T) {
	store := storage.NewMemoryStore()

	port := &entity.Port{
		ID:       "PORT1",
		Name:     "Test Port",
		City:     "Test City",
		Country:  "Test Country",
		Province: "Test Province",
	}

	store.Save(port)
	retrievedPort, found := store.Get("PORT1")

	assert.True(t, found, "Port should be found in storage")
	assert.Equal(t, port.ID, retrievedPort.ID)
	assert.Equal(t, port.Name, retrievedPort.Name)
	assert.Equal(t, port.City, retrievedPort.City)
	assert.Equal(t, port.Country, retrievedPort.Country)
	assert.Equal(t, port.Province, retrievedPort.Province)
}

func TestMemoryStore_GetNonExistentPort(t *testing.T) {
	store := storage.NewMemoryStore()

	_, found := store.Get("INVALID_PORT")
	assert.False(t, found, "Port should not be found in storage")
}

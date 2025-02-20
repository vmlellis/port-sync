package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
	"github.com/vmlellis/port-sync/src/internal/domain/service"
)

func TestSavePort(t *testing.T) {
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)

	port := &entity.Port{
		ID:       "PORT1",
		Name:     "Test Port",
		City:     "Test City",
		Country:  "Test Country",
		Province: "Test Province",
	}

	portService.SavePort(port)
	retrievedPort, found := portService.GetPort("PORT1")

	assert.True(t, found, "Port should be found in storage")
	assert.Equal(t, port.ID, retrievedPort.ID)
	assert.Equal(t, port.Name, retrievedPort.Name)
	assert.Equal(t, port.City, retrievedPort.City)
	assert.Equal(t, port.Country, retrievedPort.Country)
	assert.Equal(t, port.Province, retrievedPort.Province)
}

func TestGetNonExistentPort(t *testing.T) {
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)

	_, found := portService.GetPort("INVALID_PORT")
	assert.False(t, found, "Port should not be found in storage")
}

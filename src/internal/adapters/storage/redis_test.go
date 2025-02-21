package storage_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

func TestRedisStore_SaveAndGet_WithMock(t *testing.T) {
	db, mock := redismock.NewClientMock()
	store := &storage.RedisStore{Client: db}

	testPort := &entity.Port{
		ID:       "PORT123",
		Name:     "Test Port",
		City:     "Test City",
		Country:  "Test Country",
		Timezone: "UTC+0",
	}

	jsonData, _ := json.Marshal(testPort)
	mock.ExpectSet("PORT123", jsonData, 0).SetVal("OK")
	mock.ExpectGet("PORT123").SetVal(string(jsonData))

	err := store.Save(context.Background(), testPort)
	assert.NoError(t, err)

	retrievedPort, found := store.Get(context.Background(), "PORT123")
	assert.True(t, found)
	assert.Equal(t, testPort.ID, retrievedPort.ID)
	assert.Equal(t, testPort.Name, retrievedPort.Name)
	assert.Equal(t, testPort.City, retrievedPort.City)
	assert.Equal(t, testPort.Country, retrievedPort.Country)
	assert.Equal(t, testPort.Timezone, retrievedPort.Timezone)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisStore_Ping_WithMock(t *testing.T) {
	db, mock := redismock.NewClientMock()
	store := &storage.RedisStore{Client: db}

	mock.ExpectPing().SetVal("PONG")

	err := store.Ping(context.Background())
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

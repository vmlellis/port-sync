package storage

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// RedisStore implements the PortRepository interface using Redis.
type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore initializes a new RedisStore.
func NewRedisStore(addr string, password string, db int) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStore{
		Client: client,
	}
}

// Save inserts or updates a port record in Redis.
func (s *RedisStore) Save(ctx context.Context, port *entity.Port) error {
	jsonData, err := json.Marshal(port)
	if err != nil {
		return err
	}
	return s.Client.Set(ctx, port.ID, jsonData, 0).Err()
}

// Get retrieves a port by its unique identifier.
func (s *RedisStore) Get(ctx context.Context, id string) (*entity.Port, bool) {
	jsonData, err := s.Client.Get(ctx, id).Bytes()
	if err == redis.Nil {
		return nil, false
	} else if err != nil {
		return nil, false
	}

	var port entity.Port
	if err := json.Unmarshal(jsonData, &port); err != nil {
		return nil, false
	}
	return &port, true
}

// Ping checks the Redis connection.
func (s *RedisStore) Ping(ctx context.Context) error {
	_, err := s.Client.Ping(ctx).Result()
	return err
}

// Close closes the Redis client connection.
func (s *RedisStore) Close() error {
	return s.Client.Close()
}

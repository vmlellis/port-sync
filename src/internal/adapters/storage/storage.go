package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
)

type StorageOpts struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// NewPortRepository initializes the appropriate storage based on the config.
func NewPortRepository(ctx context.Context, storageType string, opts StorageOpts) (contract.PortRepository, error) {
	storageMap := map[string]contract.PortRepository{
		"internal": NewMemoryStore(),
		"redis":    NewRedisStore(opts.RedisAddr, opts.RedisPassword, opts.RedisDB),
	}

	storage := storageMap[storageType]
	if storage == nil {
		return nil, errors.New("storage not defined")
	}

	if err := storage.Ping(ctx); err != nil {
		return nil, fmt.Errorf("cannot connect to storage: %s", err)
	}

	return storage, nil
}

package storage_client

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
)

type StorageClient struct {
	Client *storage.Client
	Sync   sync.Once
}

func NewClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

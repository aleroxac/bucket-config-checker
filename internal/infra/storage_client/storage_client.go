package storage_client

import "context"

type StorageClientRepository interface {
	NewClient(ctx context.Context) (*interface{}, error)
}

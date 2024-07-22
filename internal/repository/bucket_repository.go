package repository

import (
	"context"

	"github.com/aleroxac/b2c/internal/dto"
)

type BucketRespository interface {
	GetBucketDetails(ctx context.Context, ch <-chan any, project string, bucket string) (interface{}, error)
	FilterBucketDetails(ctx context.Context, ch <-chan any, buckets interface{}) (interface{}, error)
	ParseBucketDetails(ctx context.Context, ch <-chan any, buckets interface{}) ([]dto.Bucket, error)

	ScanBucket(ctx context.Context, ch <-chan any, buckets []dto.Bucket) ([]interface{}, error)
	FilterBucketScan(ctx context.Context, ch <-chan any, buckets []interface{}) ([]interface{}, error)
	ParseBucketScan(ctx context.Context, ch <-chan any, buckets []interface{}) ([]dto.Bucket, error)
}

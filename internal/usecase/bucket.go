package usecase

import (
	"context"
	"fmt"
	"log"
	"slices"

	"cloud.google.com/go/storage"
	"github.com/aleroxac/b2c/internal/dto"
	"github.com/aleroxac/b2c/internal/entity"
	"github.com/aleroxac/b2c/internal/infra/storage_client"
	"google.golang.org/api/iterator"
)

func getPath(project string, bucket string) string {
	return fmt.Sprintf("%s/%s", project, bucket)
}

func getRules(lifecycle_rules []storage.LifecycleRule) []dto.Rule {
	var rules []dto.Rule
	for _, r := range lifecycle_rules {
		rule := dto.Rule{
			Action: dto.Action{
				Type: r.Action.Type,
			},
			Condition: dto.Condition{
				IsLive:           int(r.Condition.Liveness),
				NumNewerVersions: int(r.Condition.NumNewerVersions),
			},
		}
		rules = append(rules, rule)
	}
	return rules
}

func getLifecycle(rules []dto.Rule) dto.LifeCycleConfig {
	return dto.LifeCycleConfig{
		Rules: rules,
	}
}

func getLocation(location string, location_type string) dto.Location {
	return dto.Location{
		Type: location_type,
		Name: location,
	}
}

func getPublic(client *storage.Client, ctx context.Context, bucket string, prevention_type storage.PublicAccessPrevention) dto.PublicAccessConfig {
	isPublic := false
	policy, _ := client.Bucket(bucket).IAM().Policy(ctx)
	if slices.Contains(policy.Members("roles/storage.objectViewer"), "allUsers") {
		isPublic = true
	}

	public := dto.PublicAccessConfig{
		IsPublic:       isPublic,
		PreventionType: prevention_type.String(),
	}

	return public
}

func getLogging(logging *storage.BucketLogging) dto.LoggingConfig {
	return dto.LoggingConfig(*logging)
}

func GetDetails(bucket string, project string) ([]dto.Bucket, error) {
	ctx_client := context.Background()
	defer ctx_client.Done()
	client, err := storage_client.NewClient(ctx_client)
	if err != nil {
		log.Fatalf("Fail to create a new storage client: %v", err)
	}

	ctx_buckets := context.Background()
	ctx_buckets.Done()
	buckets_iterator := client.Buckets(ctx_buckets, project)
	var buckets []dto.Bucket

	for {
		bucket_attrs, err := buckets_iterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Fail to get bucket atributes: %v", err)
		}

		path := getPath(project, bucket_attrs.Name)
		rules := getRules(bucket_attrs.Lifecycle.Rules)
		lifecycle := getLifecycle(rules)
		location := getLocation(bucket_attrs.LocationType, bucket_attrs.Location)
		public := getPublic(client, ctx_buckets, bucket_attrs.Name, bucket_attrs.PublicAccessPrevention)
		logging := getLogging(bucket_attrs.Logging)

		bucket := entity.NewBucket(path,
			bucket_attrs.Name,
			project,
			bucket_attrs.StorageClass,
			lifecycle,
			location,
			logging,
			public,
			bucket_attrs.VersioningEnabled,
		)

		buckets = append(buckets, *bucket)
	}

	return buckets, nil
}

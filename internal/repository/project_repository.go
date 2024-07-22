package repository

import (
	"context"

	"github.com/aleroxac/b2c/internal/dto"
)

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]string, error)

	FilterProjects(ctx context.Context, ch <-chan any, projects []interface{}) ([]interface{}, error)
	ParseProjects(ctx context.Context, ch <-chan any, projects []interface{}) ([]dto.Project, error)

	GetProjectBuckets(ctx context.Context, ch <-chan any, project dto.Project) ([]interface{}, error)
	FilterProjectBuckets(ctx context.Context, ch <-chan any, project dto.Project) ([]interface{}, error)
	ParseProjectBuckets(ctx context.Context, ch <-chan any, project dto.Project) (dto.Project, error)
}

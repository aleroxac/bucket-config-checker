package entity

import "github.com/aleroxac/b2c/internal/dto"

func NewProject(name string, buckets []dto.Bucket) *dto.Project {
	return &dto.Project{
		Name:    name,
		Buckets: buckets,
	}
}

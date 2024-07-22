package entity

import "github.com/aleroxac/b2c/internal/dto"

func NewBucket(
	path string,
	name string,
	project string,
	class string,
	lifecycle dto.LifeCycleConfig,
	location dto.Location,
	logging dto.LoggingConfig,
	public dto.PublicAccessConfig,
	versioning bool) *dto.Bucket {

	return &dto.Bucket{
		Path:       path,
		Name:       name,
		Project:    project,
		Class:      class,
		LifeCycle:  lifecycle,
		Location:   location,
		Logging:    logging,
		Public:     public,
		Versioning: versioning,
	}
}

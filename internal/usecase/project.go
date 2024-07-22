package usecase

import (
	"context"
	"log"

	"github.com/aleroxac/b2c/cfg"
)

func GetProjects(ctx context.Context) ([]cfg.Project, error) {
	config, err := cfg.Load()
	if err != nil {
		log.Fatalf("Fail to load config: %v", err)
		return nil, err
	}

	var projects []cfg.Project
	for _, provider := range config.Providers {
		if provider.Name == "gcp" {
			projects = provider.Projects
		}
	}

	return projects, nil
}

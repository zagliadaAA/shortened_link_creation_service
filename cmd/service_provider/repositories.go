package service_provider

import (
	"context"

	"shortened_link_creation_service/internal/adapter"
	"shortened_link_creation_service/internal/adapter/in_memory_repository"
	"shortened_link_creation_service/internal/adapter/postgres_repository"
)

func (sp *ServiceProvider) GetRepository() adapter.Repository {
	if sp.repo == nil {
		if sp.repoType == "postgresRepo" {
			sp.GetPostgresRepository()
		} else {
			sp.GetInMemoryRepository()
		}
	}

	return sp.repo
}

func (sp *ServiceProvider) GetInMemoryRepository() {
	if sp.repo == nil {
		sp.repo = in_memory_repository.NewInMemoryRepository()
	}
}

func (sp *ServiceProvider) GetPostgresRepository() {
	if sp.repo == nil {
		sp.repo = postgres_repository.NewPostgresRepo(sp.GetDbCluster(context.Background()))
	}
}

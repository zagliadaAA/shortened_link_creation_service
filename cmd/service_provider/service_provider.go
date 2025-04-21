package service_provider

import (
	"shortened_link_creation_service/internal/adapter"
	"shortened_link_creation_service/internal/adapter/postgres_repository/config"
	"shortened_link_creation_service/internal/controller/link_controller"
	"shortened_link_creation_service/internal/usecase/link_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	linkUseCase *link_usecase.UseCase

	repo     adapter.Repository
	repoType string

	linkController *link_controller.Controller
}

func NewServiceProvider(repoType string) *ServiceProvider {
	return &ServiceProvider{
		repoType: repoType,
	}
}

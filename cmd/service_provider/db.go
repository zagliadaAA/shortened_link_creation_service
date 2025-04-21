package service_provider

import (
	"context"
	"log"

	"shortened_link_creation_service/internal/adapter/postgres_repository/config"
)

func (sp *ServiceProvider) GetDbCluster(ctx context.Context) *config.Cluster {
	if sp.dbCluster == nil {
		dbCluster, err := config.NewCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}

		sp.dbCluster = dbCluster
	}

	return sp.dbCluster
}

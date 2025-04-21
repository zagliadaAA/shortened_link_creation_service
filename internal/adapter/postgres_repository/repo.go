package postgres_repository

import (
	"shortened_link_creation_service/internal/adapter/postgres_repository/config"
)

type PostgresRepo struct {
	cluster *config.Cluster
}

func NewPostgresRepo(cluster *config.Cluster) *PostgresRepo {
	return &PostgresRepo{
		cluster: cluster,
	}
}

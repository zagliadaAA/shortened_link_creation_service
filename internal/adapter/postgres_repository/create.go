package postgres_repository

import (
	"shortened_link_creation_service/internal/domain"
)

func (pr *PostgresRepo) Create(link *domain.Link) (*domain.Link, error) {

	return link, nil
}

package postgres_repository

import (
	"context"
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (pr *PostgresRepo) Create(link *domain.Link) (*domain.Link, error) {
	query := "INSERT INTO links(url, short_url) VALUES ($1, $2) RETURNING id;"

	err := pr.cluster.Conn.QueryRow(context.Background(), query, link.URL, link.ShortURL).Scan(&link.ID)
	if err != nil {
		return nil, fmt.Errorf("create: ошибка при обращении к БД: %w", err)
	}

	return link, nil
}

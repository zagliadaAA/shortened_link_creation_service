package postgres_repository

import (
	"context"
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (pr *PostgresRepo) FindByURL(url string) (*domain.Link, error) {
	query := "SELECT id, url, short_url FROM links WHERE url = $1;"

	var link domain.Link

	err := pr.cluster.Conn.QueryRow(context.Background(), query, url).Scan(&link.ID, &link.URL, &link.ShortURL)
	if err != nil {
		return nil, fmt.Errorf("FindByURL: ошибка при обращении к БД: %w", err)
	}

	return &link, nil
}

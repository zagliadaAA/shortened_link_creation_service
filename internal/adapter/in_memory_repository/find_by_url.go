package in_memory_repository

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (mr *InMemoryRepo) FindByURL(url string) (*domain.Link, error) {
	for _, link := range mr.linkMap {
		if link.URL == url {
			return link, nil
		}
	}

	return nil, fmt.Errorf("URL не найден")
}

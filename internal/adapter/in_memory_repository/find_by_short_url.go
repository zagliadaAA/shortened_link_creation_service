package in_memory_repository

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (mr *InMemoryRepo) FindByShortURL(shortURL string) (*domain.Link, error) {
	for _, link := range mr.linkMap {
		if link.ShortURL == shortURL {
			return link, nil
		}
	}

	return nil, fmt.Errorf("короткий URL не найден")
}

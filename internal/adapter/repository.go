package adapter

import (
	"shortened_link_creation_service/internal/domain"
)

type Repository interface {
	Create(link *domain.Link) (*domain.Link, error)
	FindByURL(url string) (*domain.Link, error)
	FindByShortURL(shortURL string) (*domain.Link, error)
}

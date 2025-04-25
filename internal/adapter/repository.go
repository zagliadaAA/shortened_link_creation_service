package adapter

import (
	"shortened_link_creation_service/internal/domain"
)

type Repository interface {
	Create(link *domain.Link) (*domain.Link, error)
	GetLinkByURL(url string) (*domain.Link, error)
	GetLinkByShortURL(shortURL string) (*domain.Link, error)
}

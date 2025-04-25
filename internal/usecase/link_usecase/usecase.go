package link_usecase

import (
	"shortened_link_creation_service/internal/domain"
)

type UseCase struct {
	repo repo
}

type repo interface {
	GetLinkByURL(url string) (*domain.Link, error)
	Create(link *domain.Link) (*domain.Link, error)
	GetLinkByShortURL(shortURL string) (*domain.Link, error)
}

func NewUseCase(repo repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

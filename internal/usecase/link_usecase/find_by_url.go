package link_usecase

import (
	"shortened_link_creation_service/internal/domain"
)

func (uc *UseCase) FindByURL(url string) (*domain.Link, error) {
	return &domain.Link{
		URL:      url,
		ShortURL: "123456",
	}, nil
}

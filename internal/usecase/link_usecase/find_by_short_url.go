package link_usecase

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (uc *UseCase) FindByShortURL(shortURL string) (*domain.Link, error) {
	link, err := uc.repo.FindByShortURL(shortURL)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.FindByShortURL: %w", err)
	}

	return link, nil
}

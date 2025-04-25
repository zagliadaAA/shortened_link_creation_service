package link_usecase

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (uc *UseCase) GetLinkByShortURL(shortURL string) (*domain.Link, error) {
	link, err := uc.repo.GetLinkByShortURL(shortURL)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.GetLinkByShortURL: %w", err)
	}

	return link, nil
}

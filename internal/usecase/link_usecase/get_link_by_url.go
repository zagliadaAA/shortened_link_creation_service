package link_usecase

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (uc *UseCase) GetLinkByURL(url string) (*domain.Link, error) {
	link, err := uc.repo.GetLinkByURL(url)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.GetLinkByURL: %w", err)
	}

	return link, nil
}

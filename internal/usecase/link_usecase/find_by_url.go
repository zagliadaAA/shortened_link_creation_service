package link_usecase

import (
	"fmt"

	"shortened_link_creation_service/internal/domain"
)

func (uc *UseCase) FindByURL(url string) (*domain.Link, error) {
	link, err := uc.repo.FindByURL(url)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.FindByURL: %w", err)
	}

	return link, nil
}

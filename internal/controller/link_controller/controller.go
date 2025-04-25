package link_controller

import (
	"shortened_link_creation_service/internal/domain"
)

type Controller struct {
	linkUseCase linkUseCase
}

type linkUseCase interface {
	GetLinkByURL(url string) (*domain.Link, error)
	CreateShortURL(URL string) (*domain.Link, error)
	GetLinkByShortURL(shortURL string) (*domain.Link, error)
}

func NewController(useCase linkUseCase) *Controller {
	return &Controller{
		linkUseCase: useCase,
	}
}

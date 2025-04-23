package link_controller

import (
	"shortened_link_creation_service/internal/domain"
)

type Controller struct {
	linkUseCase linkUseCase
}

type linkUseCase interface {
	FindByURL(url string) (*domain.Link, error)
	CreateShortURL(URL string) (*domain.Link, error)
	FindByShortURL(shortURL string) (*domain.Link, error)
}

func NewController(useCase linkUseCase) *Controller {
	return &Controller{
		linkUseCase: useCase,
	}
}

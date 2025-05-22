package grpc

import (
	"shortened_link_creation_service/internal/domain"
	pb "shortened_link_creation_service/protos/gen/proto"
)

type Controller struct {
	linkUseCase linkUseCase
	pb.UnimplementedShortenerURLServer
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

package service_provider

import (
	"shortened_link_creation_service/internal/controller/grpc"
)

func (sp *ServiceProvider) GetLinkController() *grpc.Controller {
	if sp.linkController == nil {
		sp.linkController = grpc.NewController(sp.GetLinkUseCase())
	}

	return sp.linkController
}

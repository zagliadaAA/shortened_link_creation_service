package service_provider

import (
	"shortened_link_creation_service/internal/controller/link_controller"
)

func (sp *ServiceProvider) GetLinkController() *link_controller.Controller {
	if sp.linkController == nil {
		sp.linkController = link_controller.NewController(sp.GetLinkUseCase())
	}

	return sp.linkController
}

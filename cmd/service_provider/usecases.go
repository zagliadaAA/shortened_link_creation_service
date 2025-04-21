package service_provider

import (
	"shortened_link_creation_service/internal/usecase/link_usecase"
)

func (sp *ServiceProvider) GetLinkUseCase() *link_usecase.UseCase {
	if sp.linkUseCase == nil {
		sp.linkUseCase = link_usecase.NewUseCase(sp.GetRepository())
	}

	return sp.linkUseCase
}

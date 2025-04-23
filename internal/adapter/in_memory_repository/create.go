package in_memory_repository

import (
	"shortened_link_creation_service/internal/domain"
)

func (mr *InMemoryRepo) Create(link *domain.Link) (*domain.Link, error) {
	id := mr.getNextIdentifier()
	link.SetID(id)

	mr.linkMap[link.ID] = link

	return link, nil
}

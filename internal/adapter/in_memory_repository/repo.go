package in_memory_repository

import (
	"shortened_link_creation_service/internal/domain"
)

type InMemoryRepo struct {
	linkMap map[int]*domain.Link
	id      int
}

func NewInMemoryRepository() *InMemoryRepo {
	return &InMemoryRepo{
		linkMap: make(map[int]*domain.Link),
		id:      1,
	}
}

func (mr *InMemoryRepo) getNextIdentifier() int {
	currentID := mr.id
	mr.id++

	return currentID
}

package services

import (
	"github.com/atastrophic/memory-service/internal/datastore"
	"github.com/atastrophic/memory-service/internal/models/domain"
)

type MemoryService interface{}

func GetmemoryService(store datastore.WarmStore) MemoryService {
	return &_MemoryService{
		store: store,
	}
}

type _MemoryService struct {
	store datastore.WarmStore
}

func (service _MemoryService) GetMemories() []domain.Memory {
	return []domain.Memory{}
}

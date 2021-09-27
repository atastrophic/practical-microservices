package datastore

import "github.com/atastrophic/memory-service/internal/models/domain"

type _InMemory struct {
	storage map[string]domain.Memory
}

func GetInMemoryStore() WarmStore {
	return _InMemory{
		storage: make(map[string]domain.Memory),
	}
}

func (store _InMemory) GetMemories() ([]domain.Memory, error) {
	memories := make([]domain.Memory, len(store.storage))

	for _, memory := range store.storage {
		memories = append(memories, memory)
	}

	return memories, nil
}

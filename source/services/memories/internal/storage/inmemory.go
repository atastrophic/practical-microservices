package storage

type InMemoryStore struct {
	store map[string]string
}

func GetInMemoryStore() HotStorage {
	return InMemoryStore{}
}

func (s InMemoryStore) GetMemories() ([]string, error) {
	memories := []string{}

	for _, memory := range s.store {
		memories = append(memories, memory)
	}

	return memories, nil
}

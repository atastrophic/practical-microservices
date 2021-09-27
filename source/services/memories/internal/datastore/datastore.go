package datastore

import "github.com/atastrophic/memory-service/internal/models/domain"

type WarmStore interface {
	GetMemories() ([]domain.Memory, error)
}

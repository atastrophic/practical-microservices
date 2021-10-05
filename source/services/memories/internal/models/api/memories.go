package api

import "github.com/atastrophic/practical-microservices/internal/models/domain"

type MemoriesResponse struct {
	Data []domain.Memory `json:"data"`
}

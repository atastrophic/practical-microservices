package domain

import "github.com/google/uuid"

type Memory struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

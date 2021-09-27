package domain

import "github.com/google/uuid"

type Memory struct {
	Id   uuid.UUID
	text string
}

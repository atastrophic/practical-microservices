package storage

type HotStorage interface {
	GetMemories() ([]string, error)
}

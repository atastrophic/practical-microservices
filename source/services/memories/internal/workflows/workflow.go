package workflows

import (
	"encoding/json"

	"github.com/atastrophic/practical-microservices/internal/models/domain"
	"github.com/atastrophic/practical-microservices/internal/storage"
)

type Workflow struct {
	hotstore storage.HotStorage
}

func GetWorkflow(hotstore storage.HotStorage) Workflow {
	return Workflow{
		hotstore: hotstore,
	}
}

func (workflow Workflow) GetMemories() ([]domain.Memory, error) {
	memorystrs, err := workflow.hotstore.GetMemories()

	if err != nil {
		return nil, err
	}

	memories := []domain.Memory{}

	for _, m := range memorystrs {

		var memory domain.Memory
		err := json.Unmarshal([]byte(m), &memory)

		if err != nil {
			return nil, err
		}

		memories = append(memories, memory)
	}

	return memories, nil
}

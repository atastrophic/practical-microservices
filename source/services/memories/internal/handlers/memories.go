package handlers

import (
	"fmt"
	"net/http"

	"github.com/atastrophic/memory-service/internal/services"
	"github.com/labstack/echo"
)

func GetMemoriesHandler(service services.MemoryService) MemoriesHandler {
	return &_MemoriesHandler{
		service: service,
	}
}

type MemoriesHandler interface {
	GetMemories(c echo.Context) error
	GetMemory(c echo.Context) error
}

type _MemoriesHandler struct {
	service services.MemoryService
}

func (handler *_MemoriesHandler) GetMemories(c echo.Context) error {
	return c.String(http.StatusOK, "All memories endpoint")
}

func (handler *_MemoriesHandler) GetMemory(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%s memory here.", c.Param("resourceId")))
}

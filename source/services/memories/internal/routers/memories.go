package routers

import (
	"github.com/atastrophic/memory-service/internal/handlers"
	"github.com/atastrophic/memory-service/internal/services"
	"github.com/labstack/echo"
)

type MemoriesRouter interface {
	RegisterRoutes(e *echo.Group)
}

type _MemoriesRouter struct {
	handler handlers.MemoriesHandler
}

func TheMemoryRouter(service services.MemoryService) MemoriesRouter {
	return &_MemoriesRouter{
		handler: handlers.GetMemoriesHandler(service),
	}
}

func (router *_MemoriesRouter) RegisterRoutes(g *echo.Group) {
	memories := g.Group("/memories")

	memories.GET("/", router.handler.GetMemories)
	memories.GET("/:resourceId", router.handler.GetMemory)
}

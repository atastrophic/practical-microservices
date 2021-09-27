package routers

import (
	"github.com/atastrophic/memory-service/internal/services"
	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo, service services.MemoryService) {

	apiGroup := e.Group("/api")

	memoriesRouter := TheMemoryRouter(service)
	memoriesRouter.RegisterRoutes(apiGroup)
}

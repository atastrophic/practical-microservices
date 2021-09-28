package routes

import (
	"github.com/atastrophic/practical-microservices/internal/handlers"
	"github.com/atastrophic/practical-microservices/internal/workflows"
	"github.com/labstack/echo"
)

type Router struct {
	handler handlers.Handler
}

func GetRouter(workflow workflows.Workflow) Router {
	handler := handlers.GetHandler(workflow)
	return Router{
		handler: handler,
	}
}

func (router Router) RegisterRoutes(e *echo.Echo) {

	memories := e.Group("/memories")
	memories.GET("/", router.handler.GetMemories)
	memories.GET("", router.handler.GetMemories)

	memories.GET("/:memoryId", router.handler.GetMemory)             // get specific memroy
	memories.PUT("/:memoryId", router.handler.CreateOrReplaceMemory) // Replace or Creeate
	memories.POST("/", router.handler.CreateMemory)                  // Creeate a memory
	memories.POST("", router.handler.CreateMemory)                   // create a memory
	memories.DELETE("/:memoryId", router.handler.DeleteMemory)       // Delete a memory & respective resources i.e. photos, comments
	memories.PATCH("/:memoryId", router.handler.UpdateMemory)        // Update Memory
}

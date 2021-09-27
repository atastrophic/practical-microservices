package main

import (
	"github.com/atastrophic/memory-service/internal/datastore"
	"github.com/atastrophic/memory-service/internal/routers"
	"github.com/atastrophic/memory-service/internal/services"
	"github.com/labstack/echo"
)

var (
	service services.MemoryService
)

func main() {

	e := echo.New()

	service = services.GetmemoryService(
		datastore.GetInMemoryStore(),
	)

	// Register routes.
	routers.RegisterRoutes(e, service)

	e.Logger.Fatal(e.Start(":5000"))
}

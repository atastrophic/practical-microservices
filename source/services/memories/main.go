package main

import (
	"net/http"

	"github.com/atastrophic/practical-microservices/internal/routes"
	"github.com/atastrophic/practical-microservices/internal/storage"
	"github.com/atastrophic/practical-microservices/internal/workflows"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome.")
	})

	hotstore := storage.GetInMemoryStore()
	service := workflows.GetWorkflow(hotstore)
	router := routes.GetRouter(service)
	router.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":5000"))
}

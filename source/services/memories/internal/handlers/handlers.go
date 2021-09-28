package handlers

import (
	"net/http"

	"github.com/atastrophic/practical-microservices/internal/workflows"
	"github.com/labstack/echo"
)

type Handler struct {
	workflow workflows.Workflow
}

func GetHandler(workflow workflows.Workflow) Handler {
	return Handler{
		workflow: workflow,
	}
}

func (h Handler) GetMemories(c echo.Context) error {
	memories, err := h.workflow.GetMemories()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, memories)
}

func (handler Handler) GetMemory(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Practical Microservices.")
}

func (handler Handler) CreateOrReplaceMemory(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Practical Microservices.")
}

func (handler Handler) CreateMemory(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Practical Microservices.")
}

func (handler Handler) DeleteMemory(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Practical Microservices.")
}

func (handler Handler) UpdateMemory(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Practical Microservices.")
}

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task" validate:"required,min=3"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Validation"},
	{ID: 1, Task: "Add Middleware"},
}

var validate = validator.New()

func main() {
	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})
}

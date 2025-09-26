package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Init module"},
	{ID: 2, Task: "Build API"},
}

func main() {
	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})

	e.POST("/todos", func(c echo.Context) error {
		var newTodo Todo
		if err := c.Bind(&newTodo); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Bad Request",
			})
		}

		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
		return c.JSON(http.StatusCreated, newTodo)
	})

	e.Logger.Fatal(e.Start(":8834"))
}

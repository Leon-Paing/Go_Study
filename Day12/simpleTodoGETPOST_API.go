package main

import (
	"net/http"
	"strconv"

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
			return c.JSON(http.StatusBadGateway, map[string]string{
				"error": "Not Available",
			})
		}

		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
		return c.JSON(http.StatusCreated, newTodo)
	})

	e.GET("/todo/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}
		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(http.StatusOK, todo)
			}
		}
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Data Not Found!",
		})
	})

	e.Logger.Fatal(e.Start(":8834"))
}

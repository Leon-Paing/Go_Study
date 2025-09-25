package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Echo"},
	{ID: 2, Task: "Write API"},
}

func main() {
	e := echo.New()

	e.GET("todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})

	e.GET("todo/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "500 Server Error",
			})
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(http.StatusOK, todo)
			}
		}
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Not Found",
		})
	})

	fmt.Println("Server running at http://localhost:8834")
	e.Logger.Fatal(e.Start(":8834"))
}

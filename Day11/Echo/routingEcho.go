package main

import (
	"fmt"
	"net/http"

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
		id := c.Param("id")
		return c.String(http.StatusOK, "Todo ID: "+id)
	})

	fmt.Println("Server running at http://localhost:8834")
	e.Logger.Fatal(e.Start(":8834"))
}

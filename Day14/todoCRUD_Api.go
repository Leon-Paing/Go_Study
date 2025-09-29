package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task" validate:"required,min=3"`
	Status bool   `json:"status"`
}

var todos = []Todo{
	{ID: 1, Task: "Write API", Status: true},
	{ID: 2, Task: "Add Validation", Status: false},
	{ID: 3, Task: "Test API", Status: false},
}

var validate = validator.New()

func main() {
	e := echo.New()
	// e.HTTPErrorHandler = customErrorHander

	e.GET("/todos", ListTodos)
	e.GET("/todos/:id", GetTodo)
	// e.POST("/todos", AddTodo)
	// e.PUT("/todos/:id", UpdateTodo)
	// e.DELETE("todos/:id", DeleteTodo)

	e.Logger.Fatal(e.Start(":8834"))
}

func ListTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	for _, t := range todos {
		if t.ID == id {
			return c.JSON(http.StatusOK, t)
		}
	}

	return c.JSON(http.StatusBadGateway, "Not found related Data")
}

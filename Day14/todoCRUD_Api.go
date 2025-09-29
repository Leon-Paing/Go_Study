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
	e.POST("/todos", AddTodo)
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

func AddTodo(c echo.Context) error {
	var newTodo Todo
	if err := c.Bind(&newTodo); err != nil {
		c.JSON(http.StatusBadGateway, "Invalid JSON")
	}

	if err := validate.Struct(newTodo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	newTodo.ID = len(todos) + 1
	newTodo.Status = false
	todos = append(todos, newTodo)

	return c.JSON(http.StatusCreated, newTodo)
}

func customErrorHander(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{} = "Something went wrong"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	}

	c.JSON(code, map[string]interface{}{
		"error":  message,
		"status": code,
		"path":   c.Request().URL,
		"method": c.Request().Method,
	})
}

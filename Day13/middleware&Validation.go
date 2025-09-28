package main

import (
	"net/http"
	"strconv"

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

	e.POST("/todos", func(c echo.Context) error {
		var newTodo Todo
		if err := c.Bind(&newTodo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
		}

		if err := validate.Struct(newTodo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)

		return c.JSON(http.StatusCreated, newTodo)

	})

	e.GET("/todo/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
		}

		for _, t := range todos {
			if t.ID == id {
				return c.JSON(http.StatusOK, t)
			}
		}

		return echo.NewHTTPError(http.StatusNotFound, "Todo Not Found")
	})

	e.Logger.Fatal(e.Start(":8834"))
}

func customErrorHander(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Something went wrong"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	}

	c.JSON(code, map[string]interface{}{
		"error":  message,
		"status": code,
		"path":   c.Request().RequestURI,
		"method": c.Request().Method,
	})
}

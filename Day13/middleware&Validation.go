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

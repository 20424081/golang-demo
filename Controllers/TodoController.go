package Controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type todo struct {
	Task string `json:"task"`
	Status bool `json:"status"`
}

func GetTodos(c echo.Context) error{
	todos := []todo{
		{Task: "tesst", Status: true},
	}
	return c.JSON(
		http.StatusOK,
		Resp{
			Result: todos,
			Count:  1,
		})
}

func GetTodo(c echo.Context) error{
	return c.JSON(
		http.StatusOK,
		Resp{
			Count:  1,
		})
}

func CreateTodo(c echo.Context) error{
	return c.JSON(
		http.StatusOK,
		Resp{
			Count:  1,
		})
}

func EditTodo(c echo.Context) error{
	return c.JSON(
		http.StatusOK,
		Resp{
			Count:  1,
		})
}
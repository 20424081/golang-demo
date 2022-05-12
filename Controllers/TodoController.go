package Controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"goapp/Models"
)

func GetTodos(c echo.Context) error {
	var todos []Models.ToDo
	err := Models.GetTodos(&todos, "1", "10", "0", "")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Have error",
			})
	}
	return c.JSON(
		http.StatusOK,
		Resp{
			Result: todos,
			Count:  1,
			Message: "OK",
		})
}

func GetTodo(c echo.Context) error{
	var todo Models.ToDo
	id := c.Param("id")
	err := Models.GetATodo(&todo, id, "1")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Have error",
			})
	}
	return c.JSON(
		http.StatusOK,
		Resp{
			Result: todo,
			Message: "OK",
		})
}

func CreateTodo(c echo.Context) error{
	var todo Models.ToDo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	todo.UserId = 1
	err := Models.CreateATodo(&todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		Resp{
			Error: err.Error(),
			Message: "Have Error",
		})
	}
	return c.JSON(
		http.StatusCreated,
		Resp{
			Result: todo,
			Message: "Create successfully",
		})
}

func EditTodo(c echo.Context) error{
	var todo Models.ToDo
	id := c.Param("id")
	err := Models.GetATodo(&todo, id, "1")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Have error",
			})
	}
	if err = c.Bind(&todo); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	todo.UserId = 1
	err = Models.UpdateATodo(&todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		Resp{
			Error: err.Error(),
			Message: "Have Error",
		})
	}
	return c.JSON(
		http.StatusOK,
		Resp{
			Result: todo,
			Message: "Update successfully",
		})
}

func DeleteTodo(c echo.Context) error {
	var todo Models.ToDo
	id := c.Param("id")
	err := Models.DeleteATodo(&todo, id, "1")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Have error",
			})
	}
	return c.JSON(
		http.StatusNoContent,
		Resp{
			Message: "Delete successfully",
		})
}
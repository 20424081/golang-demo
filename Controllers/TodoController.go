package Controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
	"goapp/Models"
	"goapp/Utils"

)
func GetTodos(c echo.Context) (err error) {
	var todos []Models.ToDo
	var count int64
	userId := Utils.GetUserID(c.Get("user").(*jwt.Token))
	var filter Filter
	if err := c.Bind(&filter); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Have error",
			})
	}
	err = Models.GetTodos(&todos, Utils.UnitToString(userId), filter.Limit, filter.Offset, filter.Search)
	count, err = Models.CountTodos(Utils.UnitToString(userId), filter.Search)
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
			Count:  count,
			Message: "OK",
		})
}

func GetTodo(c echo.Context) (err error){
	var todo Models.ToDo
	id := c.Param("id")
	userId := Utils.GetUserID(c.Get("user").(*jwt.Token))
	err = Models.GetATodo(&todo, id, Utils.UnitToString(userId))
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

func CreateTodo(c echo.Context) (err error){
	var todo Models.ToDo
	if err = c.Bind(&todo); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	if err = c.Validate(todo); err != nil {
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	userId := Utils.GetUserID(c.Get("user").(*jwt.Token))

	todo.UserId = userId
	err = Models.CreateATodo(&todo)
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

func EditTodo(c echo.Context) (err error){
	var todo Models.ToDo
	id := c.Param("id")
	userId := Utils.GetUserID(c.Get("user").(*jwt.Token))

	err = Models.GetATodo(&todo, id, Utils.UnitToString(userId))
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
	if err = c.Validate(todo); err != nil {
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	todo.UserId = userId
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

func DeleteTodo(c echo.Context) (err error) {
	var todo Models.ToDo
	id := c.Param("id")
	userId := Utils.GetUserID(c.Get("user").(*jwt.Token))
	err = Models.DeleteATodo(&todo, id, Utils.UnitToString(userId))
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
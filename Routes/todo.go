package Routes
import (
	"github.com/labstack/echo/v4"
	"goapp/Controllers"
)
func TodoRouter(group *echo.Group){
	group.GET("", Controllers.GetTodos)
	group.POST("", Controllers.CreateTodo)
	group.GET("/:id", Controllers.GetTodo)
	group.PATCH("/:id", Controllers.EditTodo)
	group.DELETE("/:id", Controllers.DeleteTodo)

}
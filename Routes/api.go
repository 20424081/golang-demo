package Routes
import (
	"github.com/labstack/echo/v4"
	"goapp/Controllers"
	// "goapp/Middlewares"
)
func APIRouter(group *echo.Group){
	group.POST("/login", Controllers.Login)
	group.POST("/register", Controllers.Register)
	
	group.GET("/todos", Controllers.GetTodos)
	group.POST("/todos", Controllers.CreateTodo)
	group.GET("/todos/:id", Controllers.GetTodo)
	group.PATCH("/todos/:id", Controllers.EditTodo)
	group.DELETE("/todos/:id", Controllers.DeleteTodo)

}
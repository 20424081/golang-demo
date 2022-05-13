package Routes
import (
	"github.com/labstack/echo/v4"
	"goapp/Controllers"
)
func AuthRouter(group *echo.Group){
	group.POST("/login", Controllers.Login)
	group.POST("/refresh", Controllers.Refresh)
	group.POST("/register", Controllers.Register)
}
package route

import (
	c "github.com/goWeb/app/controllers"
	"github.com/labstack/echo"
)

//App é uma instancia de echo
var App *echo.Echo

//init construtor em go
func init() {
	App = echo.New()

	// A página inicial da aplicação
	App.GET("/", c.Home)

	api := App.Group("/v1")
	api.POST("/usuarios", c.Inserir)

}

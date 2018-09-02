package controllers

import (
	"net/http"

	_ "github.com/goWeb/app/models"
	"github.com/labstack/echo"
)

// Home é a página inicial da mina aplicação
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

// Home é a página inicial da mina aplicação
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	if nome != "" && email != "" {

	}

	return c.JSON(http.StatusBadRequest,map[string][string]{
		"menssagem" : "Os campos precisam ser passados"
	})

}

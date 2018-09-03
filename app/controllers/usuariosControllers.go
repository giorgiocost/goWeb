package controllers

import (
	"net/http"

	m "github.com/goWeb/app/models"
	"github.com/labstack/echo"
)

// Home é a página inicial da mina aplicação
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

// Inserir insere usuário na base de dados
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuarios m.Usuarios

	usuarios.Nome = nome
	usuarios.Email = email

	if nome != "" && email != "" {
		if _, err := m.UsuariosModel.Insert(usuarios); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"menssagem": "Não foi possível adicionar o registro o banco, tente novamente !",
			})
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"mensagem": "O registro foi armazenado com sucesso",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"menssagem": "Os campos precisam ser passados!",
	})

}

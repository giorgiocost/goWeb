package models

import l "github.com/goWeb/app/lib"

//Usuarios representa a tabela de usuários na base de dados
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

// UsuariosModel recebe a tabela do banco de dados
var UsuariosModel = l.Sess.Collection("usuarios")

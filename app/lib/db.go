package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "giorgio",
	Database: "membros",
}

// Sess é uma variavel que faz a conexão com o banco de dados
var Sess db.Database

func init() {
	var err error
	Sess, err = mysql.Open(config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

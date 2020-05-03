package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados - handler
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}

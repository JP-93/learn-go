package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectionDB() *sql.DB {
	conexao := "user=pedrobsouza dbname=pedrobsouza password=jopebas1993 host=localhost sslmode=disable" // informacoes de conexao
	db, err := sql.Open("postgres", conexao)                                                             // abrir conexao
	if err != nil {
		fmt.Errorf("Error, %s", err)
	}
	return db
}

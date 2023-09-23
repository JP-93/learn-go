package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

type confiDB struct {
	config *sql.DB
}

func (c *confiDB) ConnectionDb() *confiDB {
	psqlConfi := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgre", psqlConfi)
	if err != nil {
		errors.New("erro connection db")
	}
	return &confiDB{
		config: conn,
	}

}

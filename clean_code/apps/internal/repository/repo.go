package repository

import (
	"database/sql"
	"fmt"
	"github.com/golang-sql/sqlexp"
	"github.com/learn-go/clean_code/apps/config"
	_ "github.com/lib/pq"
)

type PsqlClient interface {
	DB() sqlexp.Querier
}

type ConnPSQL struct {
	PsqlConn sqlexp.Querier
}

func (p *ConnPSQL) DB() sqlexp.Querier {
	return p.PsqlConn
}

var DBInstance *ConnPSQL

func (p *ConnPSQL) Conn(config config.DataBaseConfig) error {
	connString := fmt.Sprintf("host=%s; user=%s; password=%s; port=%s; dbname=%s; sslmode=%s", config.Host, config.User, config.Password, config.Port, config.DbName, config.SslMode)
	connectionDB, err := sql.Open("postgres", connString)

	if err != nil {
		return fmt.Errorf("error connection dataBase %w", err)
	}
	err = connectionDB.Ping()
	if err != nil {
		return err
	}
	defer connectionDB.Close()
	fmt.Println("success connection")
	p.PsqlConn = connectionDB

	return nil
}

func NewConnection(confDBservice config.DataBaseConfig) (*ConnPSQL, error) {
	if DBInstance == nil {
		DBInstance = &ConnPSQL{}
		err := DBInstance.Conn(confDBservice)
		if err != nil {
			return nil, err
		}
	}
	return DBInstance, nil
}

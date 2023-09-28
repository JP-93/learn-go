package repository

import (
	"fmt"
	"log"

	"github.com/learn-go/clean_code/apps/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlClient interface {
	DB() *gorm.DB
}

type ConnPSQL struct {
	PsqlConn *gorm.DB
}

func (p *ConnPSQL) DB() *gorm.DB {
	return p.PsqlConn
}

var DBInstance *ConnPSQL

func (p *ConnPSQL) Conn(config config.DataBaseConfig) error {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)
	connectionDB, err := gorm.Open(postgres.Open(connString))

	if err != nil {
		log.Fatal("Error connect database", err)
		return err
	}

	fmt.Println("success connection")
	p.PsqlConn = connectionDB

	connectionDB.AutoMigrate(&UserValue{}, &PixValue{})

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

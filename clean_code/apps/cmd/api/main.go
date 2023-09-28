package main

import (
	"fmt"
	"github.com/learn-go/clean_code/apps/config"
	"github.com/learn-go/clean_code/apps/internal/repository"
)

func main() {
	cfg := config.NewConfig()
	_, err := repository.NewConnection(cfg.DbConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("rodou")
}

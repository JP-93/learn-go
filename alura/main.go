package main

import (
	"fmt"
	"log"
	"net/http"

	"alura/route"
)

func main() {
	route.ExecuteRoute()
	log.Println("Executando")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Errorf("Erro no server %s ", err)
	}
}

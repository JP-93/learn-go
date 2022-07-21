package main

import (
	"fmt"

	"alura_api/database"
	"alura_api/handle"
)

func main() {
	database.ConnectDb()
	fmt.Println("Start sever")
	fmt.Printf("Sever Start(port=8000), route: http://localhost:8000/\n")
	handle.HandleRequest()
}

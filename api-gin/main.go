package main

import (
	"fmt"

	"api-gin/database"
	"api-gin/handler"
)

func main() {
	database.ConnectionDB()
	handler.HandlerRequest()
	fmt.Printf("Sever on port 8080, route: http://localhost:8080/alunos\n")

}

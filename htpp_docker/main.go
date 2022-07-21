package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá mundo!")
}

func main() {
	http.HandleFunc("/Olamundo", handle)
	fmt.Printf("Sever Start(port=8080), route: http://localhost:8080/Olamundo\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("server start")
	handlerRquest()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello errado"))
}

func handlerRquest() {
	c := chi.NewRouter()
	c.Get("/teste", hello)
	http.ListenAndServe(":8082", c)
}

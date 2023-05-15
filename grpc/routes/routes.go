package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"m/v2/handlers"
	"m/v2/middleware"
)

func HandlersRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/home", handlers.Home)
	r.HandleFunc("/api/todo-task", handlers.GetAlltask).Methods("GET")
	r.HandleFunc("/api/todo-task/{id}", handlers.GetById).Methods("GET")
	r.HandleFunc("/api/todo-task/{id}", handlers.DeleteById).Methods("DELETE")
	r.HandleFunc("/api/new-task", handlers.CreateTaskTodo).Methods("PUT")
	log.Fatalf(":8000", http.ListenAndServe(":8000", r))
}

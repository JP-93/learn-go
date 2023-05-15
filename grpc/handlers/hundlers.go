package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"m/v2/db"
	"m/v2/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo ao servidor GRPC")
}

func CreateTaskTodo(w http.ResponseWriter, r *http.Request) {
	var newTesk models.ListToDo
	json.NewDecoder(r.Body).Decode(&newTesk)
	db.DB.Create(&newTesk)
	json.NewEncoder(w).Encode(newTesk)
}

func GetAlltask(w http.ResponseWriter, r *http.Request) {
	var getall []models.ListToDo
	db.DB.Find(&getall)
	json.NewEncoder(w).Encode(getall)
}
func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars
	var byId models.ListToDo
	db.DB.First(&byId, id)
	json.NewEncoder(w).Encode(byId)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars
	var byId models.ListToDo
	db.DB.Delete(&byId, id)
	json.NewEncoder(w).Encode(byId)
}

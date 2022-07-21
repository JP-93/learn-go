package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"alura_api/database"
	"alura_api/middleware"
	"alura_api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func RequestPerson(w http.ResponseWriter, r *http.Request) { // função para todas os IDs
	var p []models.Person
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnEndPoint(w http.ResponseWriter, r *http.Request) { // função de retorno do ID que queremos
	vars := mux.Vars(r)
	id := vars["id"]
	var persons models.Person
	database.DB.First(&persons, id)
	json.NewEncoder(w).Encode(persons)
}

func CreateNewPerson(w http.ResponseWriter, r *http.Request) { // cria um novo dado json dentro do banco
	var newperson models.Person
	json.NewDecoder(r.Body).Decode(&newperson)
	database.DB.Create(&newperson)
	json.NewEncoder(w).Encode(newperson)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var deletePerson models.Person
	database.DB.First(&deletePerson, id)
	json.NewDecoder(r.Body).Decode(&deletePerson)
	database.DB.Delete(&deletePerson, id)
	json.NewEncoder(w).Encode(deletePerson)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var editPerson models.Person
	database.DB.First(&editPerson, id)
	json.NewDecoder(r.Body).Decode(&editPerson)
	database.DB.Save(&editPerson)
	json.NewEncoder(w).Encode(editPerson)
}

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", Home)
	r.HandleFunc("/api/persons", RequestPerson).Methods("GET")
	r.HandleFunc("/api/persons/{id}", ReturnEndPoint).Methods("GET")
	r.HandleFunc("/api/persons", CreateNewPerson).Methods("POST")
	r.HandleFunc("/api/persons", DeletePerson).Methods("DELETE")
	r.HandleFunc("/api/persons/{id}", EditPerson).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

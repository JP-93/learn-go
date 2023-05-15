package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"api-gin/database"
	"api-gin/handler"
)

func Setup() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestHello(t *testing.T) {
	r := Setup()
	r.GET("/:nome", handler.Hello)
	req, _ := http.NewRequest("GET", "/jp", nil) // funcao que recebe a requisição e direciona para a url que será feito o GET
	resp := httptest.NewRecorder()               // armazena a resposta da request, fazendo a interface do response write
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("fail get hello")
	}
}

func TestGetAllData(t *testing.T) {
	database.ConnectionDB()
	r := Setup()
	r.GET("/alunos", handler.GetAllData)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("fail get all data")
	}

}

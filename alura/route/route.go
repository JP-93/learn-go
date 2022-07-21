package route

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"alura/module"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // função para puxar os arquivos html

func index(w http.ResponseWriter, r *http.Request) { // função para ler a request e retornar
	temp.ExecuteTemplate(w, "index", module.Getproduto())
}

func new(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func insert(w http.ResponseWriter, r *http.Request) { // respopnsavel por coletar informações que serão inseridas na requisição da pagina
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("quantidade")

		precoCov, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Errorf("Conversão invalida %v", err)
		}

		qtdConv, err := strconv.Atoi(qtd)
		if err != nil {
			fmt.Errorf("Conversão invalida %v", err)
		}
		module.CreateNewProduct(nome, descricao, precoCov, qtdConv)
	}

	http.Redirect(w, r, "/", 301)

}

func delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id") // get em alguma informação que está na url
	module.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301) // fazer redirect apos a deleção
}

func edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := module.EditProduct(idProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}
func update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("quantidade")

		idCover, err := strconv.Atoi(id)
		if err != nil {
			fmt.Errorf("Conversão invalida %v", err)
		}
		prConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Errorf("Conversão invalida %v", err)
		}

		qtdConv, err := strconv.Atoi(qtd)
		if err != nil {
			fmt.Errorf("Conversão invalida %v", err)
		}
		module.UpdateProduct(idCover, nome, descricao, prConv, qtdConv)

	}
	http.Redirect(w, r, "/", 301)
}

//Responsavel por direcionar as request e executar o codigo que se comunica com o banco de dados e a logica
func ExecuteRoute() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// fazer request para a pagina para receber o templates e tratar os dados

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "https://pt.wikipedia.org/wiki/Portal:Astronomia"
	resp, error := http.Get(url) // tratar erro e fazer a request para o site
	defer resp.Body.Close()
	check(error)
	if resp.StatusCode != 200 {
		log.Fatal("Falha ao buscar a pagina, %d", resp.StatusCode, resp.Status) //  tratar erro caso a pagina estiver off ou com algum problema
	}

	doc, error := goquery.NewDocumentFromReader(resp.Body) // faz a query no templates e faz o parse, gerar o documento usando o goquery
	check(error)
	title := doc.Find("title").Text() // aqui o parse do titulo, e retorno o templates
	fmt.Println(title)

}

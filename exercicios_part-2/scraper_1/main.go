package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// fazer request para a pagina para receber o html e tratar os dados

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	url := "https://www.uol.com.br/"
	response, error := http.Get(url) // tratar erro e fazer a request para o site
	defer response.Body.Close()
	check(error)
	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode) //  tratar erro caso a pagina estiver off ou com algum problema
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	check(error)
	black, err := doc.Find("div.exchangeBarHeader__container").Html()
	check(err)
	fmt.Println(black)

}

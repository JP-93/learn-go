package main

import (
	"encoding/json"
	"fmt"
)

// mapeando json

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tag"`
}

func main() {
	// convertendo struc  para json

	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 2000.00,
		Tags:  []string{"Promoção", "Eletronico"},
	}
	p1json, _ := json.Marshal(p1) // recebe um slice de bytes, precisa ser convertido
	fmt.Println(string(p1json))

	// convertendo json para struct

	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"Tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0])
}

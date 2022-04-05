package main

import "fmt"

func main() {
	funcPorletra := map[string]map[string]float64{
		"G": {
			"Grabiela": 2000.33,
			"Guga":     4000.33,
		},
		"J": {
			"Joao": 2000.00,
		},
		"P": {
			"Pedro": 300.00,
		},
	}

	for _, funcs := range funcPorletra {
		for nome, valor := range funcs {
			fmt.Println(nome, valor)
		}
	}
}

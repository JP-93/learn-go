package main

import "fmt"

func main() {
	funcSalario := map[string]float64{ // podemos inicializar o map na declaracao de uma variavel de operador curto
		"joao":  3000.45,
		"mari":  2000.00,
		"carla": 1000.00,
	}

	for nome, salario := range funcSalario {
		fmt.Println(nome, ": R$", salario)
	}
}

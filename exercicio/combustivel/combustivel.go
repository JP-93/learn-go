package main

import (
	"fmt"
)

func flow() int {
	var fluxo int
	fmt.Scan(&fluxo)
	return fluxo
}

func main() {
	fmt.Println("Combustivél: " +
		"Gasolina [1], " +
		"Alcool [2], " +
		"Diesel [3]" +
		"Encerrar programa[4]")

	k1 := 0
	k2 := 0
	k3 := 0

	for {
		i := flow()
		if i == 1 {
			k1++
		} else if i == 2 {
			k2++
		} else if i == 3 {
			k3++
		} else if i == 4 {
			break
		} else {
			fmt.Println("Opção inválida")
		}

	}
	fmt.Printf("Gasolina: ", k1)
	fmt.Println("Alcool: ", k2)
	fmt.Println("Diesel: ", k3)
}

// jokenpo

package main

import (
	"fmt"
)

func plinha() {
	fmt.Println(" ")
}

func main() {
	fmt.Println("Jokenpo: pedra, papel ou tesoura")
	/// Jogador 1
	plinha()
	fmt.Println("Jogador 1, escolha: ")
	var ops1 string
	fmt.Scan(&ops1)
	plinha()
	// //////// Jogador2
	plinha()
	fmt.Println("Jogador 2, escolha: ")
	var ops2 string
	fmt.Scan(&ops2)
	plinha()
	// logica do programa jogador 1

	if ops1 == "pedra" && ops2 == "pedra" {
		fmt.Println("empate")
	} else if ops1 == "pedra" && ops2 == "tesoura" {
		fmt.Println("vitoria", ops1)
	} else if ops1 == "tesoura" && ops2 == "tesoura" {
		fmt.Println("empate")
	} else if ops1 == "tesoura" && ops2 == "papel" {
		fmt.Println("virtoria", ops1)
	} else if ops1 == "papel" && ops2 == "papel" {
		fmt.Println("empate")
	} else if ops1 == "papel" && ops2 == "pedra" {
		fmt.Println("vitoria", ops1)
	}
	// logica jogador 2

	if ops2 == "pedra" && ops1 == "tesoura" {
		fmt.Println("Vitoria", ops2)
	} else if ops2 == "tesoura" && ops1 == "papel" {
		fmt.Println("Vitoria", ops2)
	} else if ops2 == "papel" && ops1 == "pedra" {
		fmt.Println("Vitoria", ops2)
	}
}

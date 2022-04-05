package main

import "fmt"

var soma = func(a, b int) int { // escopo global
	return a + b
}

func main() {
	fmt.Println(soma(1, 2)) //  escopo na funcao main

	sub := func(a, b int) int {
		return a - b

	}

	fmt.Println(sub(1, 2))
}

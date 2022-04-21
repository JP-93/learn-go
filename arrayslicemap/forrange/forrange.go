package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //  array no qual o compilador realizara a contagem
	for i, numero := range numeros {
		fmt.Println(i+1, numero)
	}

	for _, numero := range numeros { // declarado dessa forma _, com underscore o indice é ignorado
		fmt.Println(numero)
	}

}

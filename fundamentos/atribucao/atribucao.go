package main

import "fmt"

func main() {
	//atribução comum
	var b byte = 3
	fmt.Println(b)
//interferencia de tipos

	i := 3
	i += 3 // atribuição de soma
	i -= 3 // atribuição de sub
	i /= 2 // atribuição de div
	i *= 2 // atribuição de mult
	i %= 4 // atribuição de modulo
	fmt.Println(i)

// multiplos valores

x, y := 1, 2 
fmt.Println(x, y)

}
package main

import "fmt"

func troca(p1, p2 int) (primeiro, segundo int) {
	primeiro = p2
	segundo = p1
	return // retorno limpo
}

func main() {
	x, y := troca(1, 2)
	fmt.Println(x, y)
}

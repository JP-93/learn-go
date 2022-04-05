package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure() //  essa variavel vai imprimir o valor da função closure, pois ela está dentro do escopo declarado primeiro
	imprimeX()
}

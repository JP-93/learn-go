package main

import "fmt"

func inc1(n int) {
	n++
}

//ponteiro é refencia por um *
func inc2(n *int) {
	*n++ // neste caso está sendo usa para desrefenciar, ou seja, ter acesso ao valor do ponteiro
}

func main() {
	n := 1
	inc1(n) // referencia por valor
	fmt.Println(n)

	inc2(&n) // referencia por endereço, assim será aplicado a logica da função
	fmt.Println(n)
}

//& usado para obter o endereço da variável

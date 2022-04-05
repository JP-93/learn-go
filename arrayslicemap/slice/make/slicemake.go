package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10) //  slice com um array de 10 elementos ja esta associado
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // slice com capacidade de array de 20 elementos, porém com retorno de 10 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // adicionado valores no array
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // adicionado um valor a mais, o array aumenta de acordo com o aumento.
	fmt.Println(s, len(s), cap(s))
}
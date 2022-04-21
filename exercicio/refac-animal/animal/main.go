package main

import "fmt"

type numero struct {
	n1 int
	n2 int
}

func (s numero) SomaValores() int {
	res := s.n1 + s.n2
	return res
}

func main() {
	valores := numero{
		n1: 1,
		n2: 2,
	}

	fmt.Println(valores.SomaValores())
}

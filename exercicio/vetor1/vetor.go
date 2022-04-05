package main

import "fmt"

func main() {

	var vetor []int
	valor := 1
	vetor = append(vetor, valor)
	for valor <= 10 {
		calc := valor * 2
		valor++
		vetor = append(vetor, calc)
	}

	for i, seq := range vetor {
		fmt.Println("N", i, " = ", seq)
	}

}

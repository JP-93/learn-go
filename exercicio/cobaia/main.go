package main

import "fmt"

type dados struct {
	C int
	R int
	S int
}

func (som dados) somaTotal() int {
	return som.C + som.S + som.R
}

func main() {

	var recebeC string
	var recebeV int

	somaC := 0
	somaR := 0
	somaS := 0

	for i := 0; i < 4; i++ {
		fmt.Scan(&recebeC)
		fmt.Scan(&recebeV)
		if recebeC == "c" {
			somaC += recebeV
		} else if recebeC == "r" {
			somaR += recebeV
		} else if recebeC == "s" {
			somaS += recebeV
		}

	}

	storage := dados{
		C: somaC,
		R: somaR,
		S: somaS,
	}
	fmt.Println(storage.C, storage.R, storage.S)
	fmt.Println(storage.somaTotal())
}

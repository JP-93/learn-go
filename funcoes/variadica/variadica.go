package main

import "fmt"

func media(numero ...float64) float64 { // a sequencia dos pontos indica a funcao variadica
	total := 0.0
	for _, num := range numero {
		total += num
	}
	return total / float64(len(numero))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.1, 7.2, 3.1))
}

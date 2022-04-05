package main

import (
	"fmt"
)

// São homogeneos(mesmo tipo) e estaticos
func main() {
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 3.5, 5.5, 10.0
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total = +notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("%.2f", media)
}

package main

import (
	"fmt"
)

func main() {
	funcio := "1"
	salario := 5.50
	horas := 44

	calculos := (salario * float64(horas))
	fmt.Println(funcio)
	fmt.Printf("%.2f", calculos)
}

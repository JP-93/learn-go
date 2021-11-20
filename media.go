package main

import "fmt"

var n1 int
var n2 int
var n3 int
var n4 int

func medias() {
 	fmt.Println("Informe as medias")
	fmt.Scan(&n1)
	fmt.Println("n1 = ", n1)

	fmt.Println(" ")
	fmt.Scan(&n2)
	fmt.Println("n2 = ", n2)

	fmt.Println(" ")
	fmt.Scan(&n3)
	fmt.Println("n3 = ", n3)

	fmt.Println(" ")
	fmt.Scan(&n4)
	fmt.Println("n4 = ", n4)

 }

func main() {
	medias()

soma := (n1 + n2 + n3 + n4)
div := (soma/4)
fmt.Println("Total de notas", soma)
if div < 5{
	fmt.Println("Reprovado")
} else {
	fmt.Println("Aprovado")}
}


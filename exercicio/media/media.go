package main

import (
	"fmt"
)

func main() {
nota1 := 3.75
nota2 := 5.0
nota3 := 8.92
nota4 := 2.45

total := (nota1 + nota2 + nota3 + nota4) / 4

fmt.Println("Média", total)

if total < 6{
	fmt.Println("Reprovado")
} else{
	fmt.Println("Aprovado")
}

}
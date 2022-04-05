package main

import "fmt"

func imprimirnotas(nota float64){
	if nota >= 7 {
		fmt.Println("aprovado com nota", nota)
	} else {
		fmt.Println("reprovato com nota", nota)
	}
}

func main(){
	imprimirnotas(3.0)
}
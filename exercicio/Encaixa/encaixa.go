package main

import "fmt"

func main() {
	arr1 := []int{562345234857, 238547554545454, 78690}
	arr2 := []int{78690}
	if arr1[2] == arr2[0] {
		fmt.Println("encaixa")
	} else {
		fmt.Println("Não  encaixa")
	}
}

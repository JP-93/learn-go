package main

import "fmt"

func main() {
	//a := [3]int{1,2,3} //array
	//s := []int{1,2,3} //slice

	//slice é uma tecnica para acessar uma parte do array, slice nao é um array

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	fmt.Println(a2, s2)
}

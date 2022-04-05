package main

import "fmt"

// nao pode usar append em array, array é imutavel

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)
	var s1 []int // forma de declara um slice vazio

	s1 = append(s1, 3, 4, 5)
	fmt.Println(a1, s1)

	s2 := make([]int, 2) // slice com make com 2 elementos
	copy(s2, s1)
	fmt.Println(s2)
}

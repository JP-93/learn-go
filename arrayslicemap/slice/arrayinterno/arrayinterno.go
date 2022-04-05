package main

import "fmt"

func main() {

	// exemplo de uso do mesmo array interno.
	a1 := make([]int, 10)
	a2 := append(a1)
	fmt.Println(a1, a2)
	a1[2] = 2
	a2[2] = 2
	fmt.Println(a1, a2)
}

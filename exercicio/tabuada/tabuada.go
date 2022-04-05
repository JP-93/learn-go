package main

import "fmt"

func main() {

	lista := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tab := 0
	for _, numero := range lista {
		fmt.Println(numero, "x", tab, "=", numero*tab)
	}

}

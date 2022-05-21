package main

import "fmt"

func QuarterOf(mes int) {
	switch {
	case mes == 1 && mes <= 3:
		fmt.Println("primeio tri")
	case mes == 4 && mes <= 6:
		fmt.Println("segundo tri")
	case mes == 7 && mes <= 9:
		fmt.Println("terceiro tri")
	case mes == 10 && mes <= 12:
		fmt.Println("qaurto tri")
	default:
		fmt.Println("invalid")
	}

}

func main() {
	QuarterOf(13)
}

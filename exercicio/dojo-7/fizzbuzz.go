package main

import "fmt"

func main() {
	Num := []int{}

	for i := 0; i <= 100; i++ {
		Num = append(Num, i)

	}

	for _, numero := range Num {
		div3 := 3
		div5 := 5

		if numero%div3 == 0 {
			fmt.Println("fizz")
		} else if numero%div5 == 0 {
			fmt.Println("Buzz")
		} else if numero%div3 == 0 || numero%div5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(numero)
		}
	}

}

package main

import "fmt"

func main() {
	var senha int
	fmt.Scan(&senha)
	fmt.Println(senha)

	for senha != 2002 {
		fmt.Println("senha invalida")
		fmt.Println("Digite a senha novamente")
		fmt.Scan(&senha)

	}

}
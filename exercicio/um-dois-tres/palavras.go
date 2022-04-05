package main

import "fmt"

func main() {

	var um string
	fmt.Scan(&um)

	var dois string
	fmt.Scan(&dois)

	var tres string
	fmt.Scan(&tres)

	//palavras := [3]string{um, dois, tres}

	if len(um) > 3 {
		fmt.Println("incorreto")
	} else {
		fmt.Println("1")
	}

	if len(dois) > 3 {
		fmt.Println("incorreto")
	} else {
		fmt.Println("2")
	}

	if len(tres) > 5 {
		fmt.Println("incorreto")
	} else {
		fmt.Println("3")
	}
}

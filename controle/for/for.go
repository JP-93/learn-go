package main

import (
	"fmt"
	"time"
)

func main() { // não existe while em go
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println(" ")

	for i := 0; i <= 20; i += 2{
		fmt.Printf("%d ", i)
	}

// aninhamento

for i := 1; i <= 10; i ++{
	if i%2 == 0{
		fmt.Println("Par ")
	} else{
		fmt.Println("Impar ")
	}
}
// laço infinito

for {
	fmt.Println("Infinito...")
	time.Sleep(time.Second * 2)
}

}
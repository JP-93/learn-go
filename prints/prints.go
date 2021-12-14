package main

import "fmt"

func main() {
	fmt.Print("Mesma linha")
	fmt.Print(" Não pula linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")
x := 3.14

sx := fmt.Sprint(x)

fmt.Println("O valor de x é = " + sx) // concatenacao com conversao
fmt.Println("O valor de x é", x) // concatecao simples 
fmt.Printf("O valor de x é %.2f.", x) // concatenacao formatada


}

//%d = para inteiros
//%f = para flutuantes 
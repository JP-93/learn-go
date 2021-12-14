package main

import (
	"fmt"
)

func cp(trab1, trab2 bool) (bool, bool, bool){
	comprartv50 := trab1 && trab2
	comprartv32 := trab1 != trab2 // ou exclusivo
	comprarsorvete := trab1 || trab2

	return comprartv50, comprartv32, comprarsorvete
}

func main() {
	tv50, tv32, sorvete := cp(true, true)
	fmt.Printf("TV50: %t, TV 32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}

/*
Obs: Go não tem operador logico OU exclusivo, podemos usar uma diferença !=
*/



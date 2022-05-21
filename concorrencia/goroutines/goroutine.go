package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: s% (interation %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1) execução em serial
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	go fale("Maria", "Entendi!!!", 10) // execusão com go Routine em paralelo
	fale("João", "Parabéns!", 5)       //execução em serial
}

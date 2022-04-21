package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //  campo anonimo que recebe o valor de outra struct como se fosse uma herança
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "f40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("a ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}

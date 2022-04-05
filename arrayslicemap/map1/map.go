package main

import "fmt"

func main() {
	// maps precisam ser inicializados

	aprovados := make(map[int]string)

	aprovados[123456] = "ana"
	aprovados[789011] = "pedro"
	aprovados[232324] = "julio"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		//fmt.Printf("%s (CPF: %d)\n", nome, cpf)
		fmt.Println(cpf)
		fmt.Println(nome)
	}

	fmt.Println(aprovados[123456])
	delete(aprovados, 123456)
	fmt.Println(aprovados[123456])
}

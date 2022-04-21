package main

import (
	"fmt"
	"strings"
)

type pessoas struct {
	nome      string
	sobrenome string
}

func (p pessoas) getNomecompleto() string {
	return p.nome + " " + p.sobrenome
}

func (al *pessoas) setNovonome(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	al.nome = partes[0]
	al.sobrenome = partes[1]
}

func main() {

	nomes := pessoas{
		nome:      "João",
		sobrenome: "Pedro",
	}

	fmt.Println(nomes.getNomecompleto())

	nomes.setNovonome("Maria Pereira")

	fmt.Println(nomes.getNomecompleto())

}

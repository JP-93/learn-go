package main

import (
	"encoding/base64"
	"fmt"
)

type dados struct {
	nome  string
	cpf   string
	senha string
}

func main() {

	var recevier dados
	fmt.Scan(&recevier.nome)
	fmt.Scan(&recevier.cpf)
	fmt.Scan(&recevier.senha)

	recevier = dados{
		nome:  recevier.nome,
		cpf:   recevier.cpf,
		senha: recevier.senha,
	}

	sec := base64.StdEncoding.EncodeToString([]byte(recevier.senha))
	fmt.Println("Nome: ", recevier.nome)
	fmt.Println("CPF:", recevier.cpf)
	fmt.Println("Senha:", sec)
}

package main

import "fmt"

type cardapio struct {
	nome  string
	item1 float64
	item2 float64
}

func (c cardapio) Somaitem() float64 {
	return c.item1 + c.item2
}

func main() {

	pedido := cardapio{
		nome:  "hot dog",
		item1: 4.00,
		item2: 4.00,
	}
	
	fmt.Printf("total do valor %.2f R$, %v", pedido.Somaitem(), pedido.nome)

}

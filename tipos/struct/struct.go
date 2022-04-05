package main

import "fmt"

type produto struct { // declarado a struct com o indicador e o tipo. Não usar virgula
	nome     string
	preco    float64
	desconto float64
}

// Método: função receiver (receptora)

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	calc := produto{"lapis", 5.00, 0.05}
	fmt.Println(calc.precoComDesconto())

}

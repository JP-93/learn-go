package main

import "fmt"

type produto struct {
	valor  float64
	qtitem int
}

func (p produto) calc() float64 {

	return float64(p.qtitem) * p.valor
}

func main() {
	insert := produto{
		valor:  5.00,
		qtitem: 5,
	}
	fmt.Printf("Total R$ %.2f", insert.calc())

	var pagar float64
	fmt.Scan(&pagar)
	res := insert.calc() - pagar
	fmt.Println(res)

}

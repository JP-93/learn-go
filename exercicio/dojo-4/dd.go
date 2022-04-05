//exercicio do DD
package main

import "fmt"

func tele(a int) string {
	dd := a
	switch {
	case dd == 61:
		return "Brasilia"
	case dd == 71:
		return "Salvador"
	case dd == 11:
		return "São Paulo"
	case dd == 21:
		return "Rio de Janeiro"
	case dd == 32:
		return "Jiz de fora"
	case dd == 19:
		return "Campinas"
	case dd == 27:
		return "Vitoria"
	case dd == 31:
		return "BH"
	default:
		return "DDD Não castrado"
	}
}

func main() {

	fmt.Println("Digite o DDD")
	var entrada int
	fmt.Scan(&entrada)
	res := tele(entrada)
	fmt.Println(res)
}

package main

import "fmt"

func notaconceito(n float64) string{
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough  // continua o switch para o proximo bloco
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"	
	default:
		return "Nota invalida"	// caso tenha uma opção inválida
	}

} 

func main() {
	res := notaconceito(1.0)
	fmt.Println(res)
}
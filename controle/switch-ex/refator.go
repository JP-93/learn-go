package main

import	"fmt"

func notasconseito(n float64) string {
	nota := n
	switch{
	case nota >= 9 && nota >= 10:
		return "A"
	case nota >= 8 && nota >= 7:
		return "B"
	case nota >= 6 && nota >= 5:
		return "C"
	case nota >= 4 && nota >= 3:
		return "D"
	case nota >= 2 && nota >= 1:
		return "E"				
	default:
		return "nota inválida"	
	}
}

func main(){
	n1 := notasconseito(1)
	fmt.Println(n1)
}
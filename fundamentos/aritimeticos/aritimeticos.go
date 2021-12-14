package main

import (
	"fmt"
	"math"
)

func main() {
a := 4
b := 2

fmt.Println("soma", a+b)
fmt.Println("sub", a-b)
fmt.Println("div", a/b)
fmt.Println("mult", a*b)
fmt.Println("modulo", a%b )	

//pacote math
c := 6
d := 3
fmt.Println("maior", math.Max(float64(c), float64(d)))
fmt.Println("menor", math.Min(float64(c), float64(d)))
}
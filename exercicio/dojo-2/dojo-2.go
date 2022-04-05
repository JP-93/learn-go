package main

import (
	"fmt"
	"time"
)

func soma( a float64, b float64)float64{
	return a + b
}

var ops1 = 4.00
var ops2 = 4.50
var ops3 = 5.00
var ops4 = 2.00
var ops5 = 1.50

func main() {

fmt.Println("1 Hot Dog   4,00 reais")
fmt.Println("2 X-salada  4,50 reais")
fmt.Println("3 X-bacon   5,00 reais")
fmt.Println("4 Torrada   2,00 reais")
fmt.Println("5 Refri     1,50 reais")

calc := soma(ops3,ops4)
time.Sleep(time.Second * 2)
fmt.Printf("Total: R$ %.2f", calc)
}
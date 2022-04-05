package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//inteiros
	fmt.Println(1, 2, 3)
	fmt.Println("Tipos", reflect.TypeOf(1))
	// 	  inteiros sem sinal
	var b byte = 255

	fmt.Println("O valor em byte é", reflect.TypeOf(b))

	// inteiros com sinal

	i1 := math.MaxInt64
	fmt.Println("O valor maximo de int é", i1)

}

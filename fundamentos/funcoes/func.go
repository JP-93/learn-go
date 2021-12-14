package main

import (
	"fmt"
)

func soma(a int, b int) int{
	return a + b

}

func print(valor int){
	fmt.Println(valor)
}

func main(){
	res := soma(2, 3)
	print(res)
}
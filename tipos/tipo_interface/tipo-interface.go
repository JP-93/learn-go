package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // interface como variavél vazia, assim podemos atribuir valores

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{} // declarando como um tipo e atribundo valores dinamicamente, enão precisa ser do mesmo tipo
	var coisa2 dinamico = "oi"
	fmt.Println(coisa2)
}

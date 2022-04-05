package main

import (
	"fmt"
)

func main() {

	i := 1

	var p *int = nil

	p = &i // endereço de memoria

	*p++ // valor do ponteiro
	i++

	fmt.Println(p , *p, i)
}

/* anotacoes sobre ponteiro

Go não é possivel ter aritimetica com ponteiros, apenas referencialos em outros espaçoes de memoria.
ponteiro é uma referencia de memoria, é referenciado com um * ao lado da vaiavel */


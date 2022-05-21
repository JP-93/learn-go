package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) { // envi dados para o channel que será armazenado o buffer, caso o buffer lote só tera espaço apos o consumo desse dado
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("done1") //  a leitura dessa informação parte do consumo do channelque ja esta completo
	ch <- 5

}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	time.Sleep(time.Second)

}

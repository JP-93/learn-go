package main

import (
	"fmt"
	"time"
)

func workers(workerId int, data chan int) {
	for d := range data {
		fmt.Printf("Worker %d received %d\n", workerId, d)
		time.Sleep(time.Second)
	}
}

func main() {

	c := make(chan int)
	qtd := 10

	for i := 0; i < qtd; i++ {
		go workers(i, c)
	}

	for i := 0; i < 100; i++ {
		c <- i
	}

}

func publish(ch chan int) {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
}

//ch := make(chan string)
//
//go func() {
//	ch <- "Olan mundo"
//	ch <- "Nova mensagem"
//}()
//
//fmt.Println(<-ch)
//fmt.Println(<-ch)

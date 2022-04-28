package main

import (
	"fmt"
	"time"
)

/*
Channel  é uma forma de comunicação entre as goroutines
*/

func mult(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)

	c <- 3 * base

	time.Sleep(3 * time.Second)

	c <- 4 * base
}

func main() {
	c := make(chan int)
	go mult(2, c)

	a, b := <-c, <-c // reebendo dados do channel
	fmt.Println(a, b)

	fmt.Println(<-c)
}
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func ReturnCalc(s Shape) {
	fmt.Println(s.Area())
}

type Retancle struct {
	w, h float64
}

type Cicle struct {
	R float64
}

func (r *Retancle) Area() float64 {
	return r.w * r.h
}

func (c *Cicle) Area() float64 {
	return math.Pi * c.R * c.R
}

func main() {
	re := &Retancle{
		w: 3,
		h: 4,
	}
	ci := &Cicle{R: 33}
	ReturnCalc(re)
	ReturnCalc(ci)
}

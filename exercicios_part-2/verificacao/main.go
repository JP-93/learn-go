package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	valid := strings.Contains(str, "( )")
	if valid == true {
		fmt.Println("Correct")
	} else {
		fmt.Println("incorrect")
	}
}

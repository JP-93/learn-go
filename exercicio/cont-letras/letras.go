package main

import (
	"fmt"
	"strings"
)

func main() {
	upper := strings.ToUpper("rWWRRR")
	fmt.Println(strings.Count(upper, "R"))
}

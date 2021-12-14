package main
import (
	"fmt"
	"time"

)
func main() {
fmt.Println("Str:", "Ban" == "Ban") // operador de igualdade de string  return true or false
fmt.Println("!=: ", 3 !=2) // operador de desigualdade
fmt.Println("<", 3 < 2)
fmt.Println(">", 3 > 2)
fmt.Println("<=", 3 <= 2)
fmt.Println(">=", 3 >= 2)

d1 := time.Unix(0, 0)
d2 := time.Unix(0, 0)
fmt.Println(d1 == d2)


}
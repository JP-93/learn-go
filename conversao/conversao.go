// em go é necessário converter explicitamente, informando ao go qual o tipo do int

package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
//

nota := 6.9
notafinal := int(nota)
fmt.Println(notafinal)

//conversao de intereiro para str
convstr := strconv.Itoa(123)
fmt.Println("teste "+ convstr)

// str para inteiro
// retornar dois valores, caso nao for informando um str valida

num, _ := strconv.Atoi("123")
println(num)

// conversao de bool

b,_ := strconv.ParseBool("true")
if b{
	fmt.Println("Verdadeiro")
}

}
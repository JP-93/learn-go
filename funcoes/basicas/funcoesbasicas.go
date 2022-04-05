package main

import "fmt"

func f1() {
	fmt.Println("F1") //  sem parametro e sem retorno
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2) //  recebe dois paramentros e não tem return
}

func f3() string {
	return "F3" // sem parametro mas ter um retorno obrigatorio
}

func f4(p1, p2, string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) // parametro otimizado e com retorno de string obrigatorio
}

func f5() (string, string) {
	return "Retorno1", "Retorno2" //  retorno multiplos
}

func main() {
	f1()
	f2("Param1", "Param2")
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}

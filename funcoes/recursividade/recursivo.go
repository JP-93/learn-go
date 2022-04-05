package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatoralAnterior, _ := fatorial(n - 1) //aqui está declarado uma variavel que recebe a funcao e o erro
		return n * fatoralAnterior, nil
	}
}

func main() {
	res, _ := fatorial(5)
	fmt.Println(res)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}

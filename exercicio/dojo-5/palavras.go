package main

import ("fmt")

func main(){

	var entrar string
	fmt.Scan(&entrar)

	countoc := 0
	countlv := 0


	for i := 0; i <= 6; i++{
		if entrar == "sim"{
			fmt.Scan(&entrar)
			countoc += 1

		}else if entrar == "não"{
			fmt.Scan(&entrar)
			countlv += 1
		}
	}

	fmt.Println("Ainda temos", countlv, "livres")

	if countlv <= 3{
		fmt.Println("Ainda pode entrar")
	}else{
		fmt.Println("Tudo ocupado")
	}

}



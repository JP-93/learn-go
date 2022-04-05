package main

import ("fmt"
		"time"
)

func fluxo(a, b, c string){
	if a == "vertebrado" && b == "mamifero" && c == "onivoro"{
		fmt.Println("homem")
	}else if a == "vertebrado" && b == "mamifero" && c == "herbivoro"{
		fmt.Println("vaca")
	}else if a == "vertebrado" && b == "ave" && c == "carnivoro"{
		fmt.Println("aguia")
	}else if a == "vertebrado" && b == "ave" && c == "onivoro"{
		fmt.Println("pombo")
	}else if a == "invertebrado" && b == "inseto" && c == "hematofago"{
		fmt.Println("pulga")
	}else if a == "invertebrado" && b == "inseto" && c == "herbivoro"{
		fmt.Println("lagarta")
	}else if a == "invertebrado" && b == "anelideo" && c == "hematofago"{
		fmt.Println("sanguessuga")
	}else if a == "invertebrado" && b == "anelideo" && c == "onivoro"{
		fmt.Println("minhoca")
	}else{
		fmt.Println("Invalido")
	}
}

func main() {
	var ops1 string
	fmt.Scan(&ops1)

	var ops2 string
	fmt.Scan(&ops2)
	
	var ops3 string
	fmt.Scan(&ops3)

	fmt.Println("Montando arvore biologica.....")
	time.Sleep(time.Second * 5)

	fluxo(ops1, ops2, ops3)
	

}
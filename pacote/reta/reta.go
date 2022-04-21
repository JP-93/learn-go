package main

import "math"

/* Iniciando uma struc, interface, constante ou fincção com letra maiuscula, refere-se a informações Publicar, visivel dentro e fora do pacote
 Iniciando uma struc, interface, constante ou fincção com letra minuscula, refere-se a informações privadas, visivel dentro do pacote apenas
por exemplo >>  Ponto - gera algo pub
pontso ou _Ponto - gera algo privado
*/

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - b.y)
	return // retorno nomeado
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

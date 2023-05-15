package metodos

import "math"

type Retangulo struct {
	Area    float64
	Largura float64
	Base    float64
}
type Circulo struct {
	Raio float64
}

func Perimetro(r Retangulo) float64 {
	return 2 * (r.Area + r.Largura)
}

func Area(r Retangulo) float64 {
	return r.Base * r.Area
}

func CalcCirculo(c Circulo) float64 {
	return math.Pi * c.Raio * c.Raio
}

package metodos

import "testing"

func TestPerimetro(t *testing.T) {
	re := Retangulo{
		Area:    10.0,
		Largura: 10.0,
	}
	got := Perimetro(re)
	want := 40.0

	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}

}
func TestArea(t *testing.T) {
	t.Run("Retangulos", func(t *testing.T) {
		re := Retangulo{
			Area: 12.0,
			Base: 6.0,
		}
		got := Area(re)
		want := 72.0

		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	})

	t.Run("circulo", func(t *testing.T) {
		c := Circulo{Raio: 10}
		got := CalcCirculo(c)
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	})

}

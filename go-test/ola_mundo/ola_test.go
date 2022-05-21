package ola_mundo

import (
	"testing"
)

func TestHello(t *testing.T) {

	verificarMensagem := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("diz olá para um pessoa", func(t *testing.T) {
		got := Hello("joão", "")
		want := "Olá, joão"
		verificarMensagem(t, got, want)
	})

	t.Run("diz 'olá mundo !' quando uma string vazia for passada", func(t *testing.T) {
		got := Hello("", "")
		want := "Olá, mundo"

		verificarMensagem(t, got, want)
	})

	//hello em espanhol

	t.Run("Olá mundo em espanhol", func(t *testing.T) {
		got := Hello("João", "Espanhol")
		want := "Hola, João"
		verificarMensagem(t, got, want)
	})

	t.Run("Olá mundo em farances", func(t *testing.T) {

		got := Hello("João", "Frances")
		want := "Bonjour, João"
		verificarMensagem(t, got, want)

	})
}

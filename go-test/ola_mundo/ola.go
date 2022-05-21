package ola_mundo

const prefixOlaBr = "Olá, "
const prefixOlaSp = "Hola, "
const prefixFra = "Bonjour, "
const espanhol = "Espanhol"
const france = "França"

func Hello(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	if idioma == espanhol {
		return prefixOlaSp + nome
	}

	if idioma == prefixFra {
		return prefixFra + nome
	}
	return prefixOlaBr + nome
}

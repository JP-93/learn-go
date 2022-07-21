package models

type Person struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
	Info string `json:"info"`
}

var Persons []Person // variavel para mocar informações para teste

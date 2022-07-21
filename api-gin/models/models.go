package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model        //adiciona valores no banco de dados
	Nome       string `json:"nome"`
	CPF        string `json:"CPF"`
	RG         string `json:"RG"`
}

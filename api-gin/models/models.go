package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model        //adiciona valores no banco de dados
	Nome       string `json:"nome" validate:"nonzero"`
	CPF        string `json:"CPF" validate:"len=11, regexp=^[0-9]*$"`
	RG         string `json:"RG" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidateData(a *Aluno) error {
	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}

package models

import "gorm.io/gorm"

type ListToDo struct {
	gorm.Model
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Lists []ListToDo

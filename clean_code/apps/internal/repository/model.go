package repository

import "gorm.io/gorm"

type UserValue struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type PixValue struct {
	gorm.Model
	ID     int    `json:"id" gorm:"primaryKey"`
	KeyPix string `json:"key_pix"`
}
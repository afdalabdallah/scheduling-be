package models

import (
	"gorm.io/gorm"
)

type Ruangan struct {
	gorm.Model
	Nomor     string `json:"nomor_ruangan"`
	Kapasitas int    `json:"kapasitas"`
	Deskripsi string `json:"deskripsi"`
}

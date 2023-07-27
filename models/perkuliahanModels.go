package models

import (
	"gorm.io/gorm"
)

type Perkuliahan struct {
	gorm.Model
	Kelas          string `json:"kelas"`
	Sesi  		   string `json:"sesi"`
	Ruangan 	   string `json:"ruangan"`
	MataKuliahId   string `json:"mata_kuliah_id"`
}

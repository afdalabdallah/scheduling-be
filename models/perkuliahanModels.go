package models

import (
	"gorm.io/gorm"
)

type Perkuliahan struct {
	gorm.Model
	Kelas        string `json:"kelas"`
	Sesi         string `json:"sesi"`
	Ruangan      string `json:"ruangan"`
	MataKuliahId uint   `json:"mata_kuliah_id" gorm:"constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	DosenId      uint   `json:"dosen_id" gorm:"constraint:OnDelete:SET NULL, OnUpdate:CASCADE"`
	Matkul       Matkul `gorm:"foreignKey:MataKuliahId"`
	Dosen        Dosen  `gorm:"foreignKey:DosenId"`
}

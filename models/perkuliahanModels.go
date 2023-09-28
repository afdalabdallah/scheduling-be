package models

import (
	"gorm.io/gorm"
)

type Perkuliahan struct {
	gorm.Model
	Kelas    string `json:"kelas"`
	Sesi     string `json:"sesi"`
	Ruangan  string `json:"ruangan"`
	MatkulID uint   `json:"mata_kuliah_id" gorm:"constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	DosenID  uint   `json:"dosen_id" gorm:"constraint:OnDelete:SET NULL, OnUpdate:CASCADE"`
	RumpunID uint   `json:"rumpun_id" gorm:"constraint:OnDelete:SET NULL, OnUpdate:CASCADE"`
	Matkul   Matkul `gorm:"foreignKey:MatkulID"`
	Dosen    Dosen  `gorm:"foreignKey:DosenID"`
	Rumpun   Rumpun `gorm:"foreignKey:RumpunID"`
}

package models

import "gorm.io/gorm"

type Matkul struct {
	gorm.Model
	Nama     string `json:"nama"`
	KodeMK   string `json:"kode_mk" gorm:"unique"`
	SKS      int    `json:"sks"`
	Tipe     string `json:"tipe"`
	Semester int    `json:"semester"`
	RumpunID uint   `json:"rumpun_id" gorm:"constraint:OnDelete:SET NULL, OnUpdate:CASCADE, type:bigint(20) unsigned"`
	Rumpun   Rumpun `gorm:"foreignKey:RumpunID"`
}

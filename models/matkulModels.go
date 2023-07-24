package models

import "gorm.io/gorm"

type Matkul struct {
	gorm.Model
	Nama     string `json:"nama"`
	KodeMK   string `json:"kode_mk"`
	SKS      int    `json:"sks"`
	Tipe     string `json:"tipe"`
	Semester int    `json:"semester"`
	RumpunID int    `json:"rumpun_id"`
}

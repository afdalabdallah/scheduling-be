package models

import (
	"gorm.io/gorm"
)

type Dosen struct {
	gorm.Model
	Nama       string `json:"nama"`
	KodeDosen  string `json:"kode_dosen"`
	Preferensi string `json:"preferensi"`
	RumpunID   int    `json:"rumpun_id"`
}

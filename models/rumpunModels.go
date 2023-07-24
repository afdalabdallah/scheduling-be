package models

import "gorm.io/gorm"

type Rumpun struct {
	gorm.Model
	Nama    string `json:"nama"`
	KodeRMK string `json:"kode_rmk"`
}

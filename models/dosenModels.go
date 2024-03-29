package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Preferensi struct {
	Hari []string `json:"hari"`
	Sesi []string `json:"sesi"`
}

func (p *Preferensi) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, p)
	}
	return errors.New("failed to unmarshal Preferensi")
}

// Value implements the driver.Valuer interface
func (p Preferensi) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type Dosen struct {
	gorm.Model
	Nama       string     `json:"nama"`
	KodeDosen  string     `json:"kode_dosen" gorm:"unique"`
	Preferensi Preferensi `json:"preferensi"`
	RumpunID   uint       `json:"rumpun_id" gorm:"constraint:OnDelete:SET NULL, OnUpdate:CASCADE"`
	Load       int        `json:"load"`
	Rumpun     Rumpun     `gorm:"foreignKey:RumpunID"`
}

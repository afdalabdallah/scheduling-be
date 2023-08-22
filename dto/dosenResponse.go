package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

type DosenResponse struct {
	ID         int        `json:"id"`
	Nama       string     `json:"nama"`
	KodeDosen  string     `json:"kode_dosen"`
	Preferensi Preferensi `json:"preferensi"`
	Rumpun     string     `json:"rumpun_id"`
	Load       int        `json:"load"`
}

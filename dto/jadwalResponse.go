package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type JSONSlice []map[string]interface{}

func (js *JSONSlice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, js)
	}
	return errors.New("failed to unmarshal JSONSlice")
}

func (js JSONSlice) Value() (driver.Value, error) {
	return json.Marshal(js)
}

type JSONArrayString []string
func (js *JSONArrayString) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, js)
	}
	return errors.New("failed to unmarshal JSONArrayString")
}

func (js JSONArrayString) Value() (driver.Value, error) {
	return json.Marshal(js)
}


type Constraint struct {
	FirstConstraint  int `json:"first_constraint"`
	SecondConstraint int `json:"second_constraint"`
	ThirdConstraint  int `json:"third_constraint"`
	FourthConstraint int `json:"fourth_constraint"`
	FifthConstraint  int `json:"fifth_constraint"`
}


func (p *Constraint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, p)
	}
	return errors.New("failed to unmarshal Preferensi")
}

// Value implements the driver.Valuer interface
func (p Constraint) Value() (driver.Value, error) {
	return json.Marshal(p)
}


type JadwalResponse struct {
	// ID	uint `json:"id"`
	// CreatedAt time.Time `json:"created_at"`
	gorm.Model
	Data               JSONSlice  `json:"data" gorm:"type:json"`
	Fitness            float64    `gorm:"type:float" json:"fitness"`
	ViolatedConstraint Constraint `gorm:"type:json" json:"violated_constraint"`
	Skpb               JSONSlice  `json:"skpb" gorm:"type:json"`
	UnwantedSesi	JSONArrayString `json:"unwanted_sesi" gorm:"type:json"`
	ListRuangan	JSONArrayString	`json:"list_ruangan" gorm:"type:json"`
}
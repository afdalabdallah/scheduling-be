package dto

type MatkulResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	KodeMK   string `json:"kode_mk"`
	SKS      int    `json:"sks"`
	Tipe     string `json:"tipe"`
	Semester int    `json:"semester"`
	Rumpun   string `json:"rumpun"`
}

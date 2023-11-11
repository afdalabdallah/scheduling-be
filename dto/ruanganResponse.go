package dto

type RuanganResponse struct {
	ID        int    `json:"id"`
	Nomor     string `json:"nomor_ruangan"`
	Kapasitas int    `json:"kapasitas"`
	Deskripsi string `json:"deskripsi"`
}

package dto

type JadwalNewResponse struct {
	KodeDosen      string     `json:"kode_dosen"`
	KodeMataKuliah string     `json:"kode_mk"`
	Kelas          string     `json:"kelas"`
	Ruangan        string     `json:"ruangan"`
	Sesi           string     `json:"sesi"`
	Preferensi     Preferensi `json:"preferensi"`
	Tipe           string     `json:"tipe"`
	Rumpun         string     `json:"rumpun"`
}

package dto

type JadwalNewResponse struct {
	KodeDosen      string     `json:"dosen"`
	KodeMataKuliah string     `json:"mata_kuliah"`
	Kelas          string     `json:"kelas"`
	Ruangan        string     `json:"ruangan"`
	Sesi           string     `json:"sesi"`
	Preferensi     Preferensi `json:"preferensi"`
	Tipe           string     `json:"tipe"`
	Rumpun         string     `json:"rumpun"`
}

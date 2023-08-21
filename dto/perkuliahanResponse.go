package dto

type PerkuliahanResponse struct {
	ID             int    `json:"id"`
	Kelas          string `json:"kelas"`
	Sesi           string `json:"sesi"`
	Ruangan        string `json:"ruangan"`
	KodeMataKuliah string `json:"kode_mk"`
	MataKuliah     string `json:"nama_mk"`
	Rumpun         string `json:"rumpun"`
	DosenNama      string `json:"nama_dosen"`
	KodeDosen      string `json:"kode_dosen"`
	Semester       int    `json:"semester"`
}

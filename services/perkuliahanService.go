package services

import (
	"github.com/afdalabdallah/backend-web/dto"
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/dosen_repository"
	"github.com/afdalabdallah/backend-web/repository/matkul_repository"
	"github.com/afdalabdallah/backend-web/repository/perkuliahan_repository"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
)

type perkuliahanService struct {
	perkuliahanRepo perkuliahan_repository.PerkuliahanRepository
	dosenRepo       dosen_repository.DosenRepository
	matkulRepo      matkul_repository.MatkulRepository
	rumpunRepo      rumpun_repository.RumpunRepository
}

type PerkuliahanService interface {
	CreatePerkuliahan(PerkuliahanData []models.Perkuliahan) (*[]models.Perkuliahan, errs.Errs)
	GetAllPerkuliahan() (*[]dto.PerkuliahanResponse, errs.Errs)
	DeletePerkuliahan(PerkuliahanID uint) (string, errs.Errs)
	UpdatePerkuliahan(PerkuliahanID uint, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanById(PerkuliahanID uint) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanFormat() (*[]dto.JadwalNewResponse, errs.Errs)
}

func NewPerkuliahanService(rumpunRepo rumpun_repository.RumpunRepository, matkulRepo matkul_repository.MatkulRepository, dosenRepo dosen_repository.DosenRepository, perkuliahanRepo perkuliahan_repository.PerkuliahanRepository) PerkuliahanService {
	return &perkuliahanService{
		perkuliahanRepo: perkuliahanRepo,
		matkulRepo:      matkulRepo,
		dosenRepo:       dosenRepo,
		rumpunRepo:      rumpunRepo,
	}
}

func (p *perkuliahanService) CreatePerkuliahan(PerkuliahanData []models.Perkuliahan) (*[]models.Perkuliahan, errs.Errs) {
	var perkuliahanCreateRespons []models.Perkuliahan

	for _, data := range PerkuliahanData {

		matkul, errMatkul := p.matkulRepo.GetMatkulById(data.MatkulID)
		if errMatkul != nil {
			return nil, errMatkul
		}
		dosen, errDosen := p.dosenRepo.GetDosenById(data.DosenID)
		if errDosen != nil {
			return nil, errDosen
		}

		Perkuliahan := models.Perkuliahan{
			Sesi:     data.Sesi,
			Kelas:    data.Kelas,
			Ruangan:  data.Ruangan,
			MatkulID: matkul.ID,
			DosenID:  dosen.ID,
			RumpunID: matkul.Rumpun.ID,
		}
		var dosenLoads int
		dosenLoads = dosen.Load + matkul.SKS
		var dosenNew models.Dosen

		dosenNew.Nama = dosen.Nama
		dosenNew.KodeDosen = dosen.KodeDosen
		dosenNew.Preferensi = dosen.Preferensi
		dosenNew.RumpunID = dosen.RumpunID
		dosenNew.Load = dosenLoads

		dosenUpdateLoad, err := p.dosenRepo.UpdateDosen(dosen.ID, dosenNew)
		PerkuliahanCreated, err := p.perkuliahanRepo.CreatePerkuliahan(Perkuliahan)
		if err != nil {
			return nil, err
		}
		println(dosenUpdateLoad)
		perkuliahanCreateRespons = append(perkuliahanCreateRespons, *PerkuliahanCreated)
	}

	return &perkuliahanCreateRespons, nil
}

func (p *perkuliahanService) GetAllPerkuliahan() (*[]dto.PerkuliahanResponse, errs.Errs) {
	perkuliahans, err := p.perkuliahanRepo.GetAllPerkuliahan()

	var perkuliahanRespons []dto.PerkuliahanResponse
	for _, perkuliahan := range perkuliahans {

		perkuliahanResponsData := dto.PerkuliahanResponse{
			ID:             int(perkuliahan.ID),
			Sesi:           perkuliahan.Sesi,
			Kelas:          perkuliahan.Kelas,
			Ruangan:        perkuliahan.Ruangan,
			KodeMataKuliah: perkuliahan.Matkul.KodeMK,
			MataKuliah:     perkuliahan.Matkul.Nama,
			DosenNama:      perkuliahan.Dosen.Nama,
			Rumpun:         perkuliahan.Rumpun.KodeRMK,
			KodeDosen:      perkuliahan.Dosen.KodeDosen,
			Semester:       perkuliahan.Matkul.Semester,
		}
		perkuliahanRespons = append(perkuliahanRespons, perkuliahanResponsData)
	}

	if err != nil {
		return nil, err
	}

	return &perkuliahanRespons, nil
}

func (p *perkuliahanService) DeletePerkuliahan(PerkuliahanID uint) (string, errs.Errs) {
	perkuliahanData, errPerkuliahan := p.perkuliahanRepo.GetPerkuliahanById(PerkuliahanID)
	if errPerkuliahan != nil {
		return "", errPerkuliahan
	}
	matkul, errMatkul := p.matkulRepo.GetMatkulById(perkuliahanData.MatkulID)
	if errMatkul != nil {
		return "", errMatkul
	}
	dosen, errDosen := p.dosenRepo.GetDosenById(perkuliahanData.DosenID)
	if errDosen != nil {
		return "", errDosen
	}

	var dosenLoads int
	
	if(dosen.Load > 0){
		dosenLoads = dosen.Load - matkul.SKS
	}else{
		dosenLoads = dosen.Load
	}
	
	var dosenNew models.Dosen

	dosenNew.Nama = dosen.Nama
	dosenNew.KodeDosen = dosen.KodeDosen
	dosenNew.Preferensi = dosen.Preferensi
	dosenNew.RumpunID = dosen.RumpunID
	dosenNew.Load = dosenLoads

	dosenUpdateLoad, errDosenUpdate := p.dosenRepo.UpdateDosen(dosen.ID, dosenNew)
	response, err := p.perkuliahanRepo.DeletePerkuliahan(PerkuliahanID)
	println(dosenUpdateLoad, errDosenUpdate)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *perkuliahanService) UpdatePerkuliahan(PerkuliahanID uint, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	Perkuliahan := models.Perkuliahan{
		Sesi:     PerkuliahanData.Sesi,
		Kelas:    PerkuliahanData.Kelas,
		Ruangan:  PerkuliahanData.Ruangan,
		MatkulID: PerkuliahanData.MatkulID,
		DosenID:  PerkuliahanData.DosenID,
		RumpunID: PerkuliahanData.RumpunID,
	}

	PerkuliahanUpdated, err := p.perkuliahanRepo.UpdatePerkuliahan(PerkuliahanID, Perkuliahan)

	if err != nil {
		return nil, err
	}

	return PerkuliahanUpdated, nil
}

func (p *perkuliahanService) GetPerkuliahanById(PerkuliahanID uint) (*models.Perkuliahan, errs.Errs) {
	PerkuliahanData, err := p.perkuliahanRepo.GetPerkuliahanById(PerkuliahanID)

	if err != nil {
		return nil, err
	}

	return PerkuliahanData, nil
}

func (p *perkuliahanService) GetPerkuliahanFormat() (*[]dto.JadwalNewResponse, errs.Errs) {
	perkuliahans, err := p.perkuliahanRepo.GetAllPerkuliahan()

	var JadwalNewResponse []dto.JadwalNewResponse
	for _, perkuliahan := range perkuliahans {

		JadwalResponse := dto.JadwalNewResponse{
			KodeDosen:      perkuliahan.Dosen.KodeDosen,
			KodeMataKuliah: perkuliahan.Matkul.KodeMK,
			Kelas:          perkuliahan.Kelas,
			Ruangan:        perkuliahan.Ruangan,
			Sesi:           perkuliahan.Sesi,
			Preferensi:     dto.Preferensi(perkuliahan.Dosen.Preferensi),
			Tipe:           perkuliahan.Matkul.Tipe,
			Rumpun:         perkuliahan.Matkul.Rumpun.KodeRMK,
		}
		JadwalNewResponse = append(JadwalNewResponse, JadwalResponse)
	}

	if err != nil {
		return nil, err
	}

	return &JadwalNewResponse, nil
}

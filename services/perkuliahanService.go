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
	CreatePerkuliahan(PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetAllPerkuliahan() (*[]dto.PerkuliahanResponse, errs.Errs)
	DeletePerkuliahan(PerkuliahanID int) (string, errs.Errs)
	UpdatePerkuliahan(PerkuliahanID int, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanById(PerkuliahanID int) (*models.Perkuliahan, errs.Errs)
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

func (p *perkuliahanService) CreatePerkuliahan(PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	matkul, errMatkul := p.matkulRepo.GetMatkulById(PerkuliahanData.MataKuliahId)
	if errMatkul != nil {
		return nil, errMatkul
	}
	dosen, errDosen := p.dosenRepo.GetDosenById(PerkuliahanData.DosenId)
	if errDosen != nil {
		return nil, errDosen
	}

	Perkuliahan := models.Perkuliahan{
		Sesi:         PerkuliahanData.Sesi,
		Kelas:        PerkuliahanData.Kelas,
		Ruangan:      PerkuliahanData.Ruangan,
		MataKuliahId: int(matkul.ID),
		DosenId:      int(dosen.ID),
	}
	var dosenLoads int
	dosenLoads = dosen.Load + matkul.SKS
	var dosenNew models.Dosen

	dosenNew.Nama = dosen.Nama
	dosenNew.KodeDosen = dosen.KodeDosen
	dosenNew.Preferensi = dosen.Preferensi
	dosenNew.RumpunID = dosen.RumpunID
	dosenNew.Load = dosenLoads

	dosenUpdateLoad, err := p.dosenRepo.UpdateDosen(int(dosen.ID), dosenNew)
	PerkuliahanCreated, err := p.perkuliahanRepo.CreatePerkuliahan(Perkuliahan)
	if err != nil {
		return nil, err
	}
	println(dosenUpdateLoad)
	return PerkuliahanCreated, nil
}

func (p *perkuliahanService) GetAllPerkuliahan() (*[]dto.PerkuliahanResponse, errs.Errs) {
	perkuliahans, err := p.perkuliahanRepo.GetAllPerkuliahan()

	var perkuliahanRespons []dto.PerkuliahanResponse
	for _, perkuliahan := range perkuliahans {
		matkul, errMatkul := p.matkulRepo.GetMatkulById(perkuliahan.MataKuliahId)
		if errMatkul != nil {
			return nil, errMatkul
		}
		dosen, errDosen := p.dosenRepo.GetDosenById(perkuliahan.DosenId)
		if errDosen != nil {
			return nil, errDosen
		}
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkul.RumpunID)
		if errRumpun != nil {
			return nil, errRumpun
		}

		perkuliahanResponsData := dto.PerkuliahanResponse{
			ID:             int(perkuliahan.ID),
			Sesi:           perkuliahan.Sesi,
			Kelas:          perkuliahan.Kelas,
			Ruangan:        perkuliahan.Ruangan,
			KodeMataKuliah: matkul.KodeMK,
			MataKuliah:     matkul.Nama,
			DosenNama:      dosen.Nama,
			Rumpun:         rumpun.KodeRMK,
			KodeDosen:      dosen.KodeDosen,
			Semester:       matkul.Semester,
		}
		perkuliahanRespons = append(perkuliahanRespons, perkuliahanResponsData)
	}

	if err != nil {
		return nil, err
	}

	return &perkuliahanRespons, nil
}

func (p *perkuliahanService) DeletePerkuliahan(PerkuliahanID int) (string, errs.Errs) {
	perkuliahanData, errPerkuliahan := p.perkuliahanRepo.GetPerkuliahanById(PerkuliahanID)
	if errPerkuliahan != nil {
		return "", errPerkuliahan
	}
	matkul, errMatkul := p.matkulRepo.GetMatkulById(perkuliahanData.MataKuliahId)
	if errMatkul != nil {
		return "", errMatkul
	}
	dosen, errDosen := p.dosenRepo.GetDosenById(perkuliahanData.DosenId)
	if errDosen != nil {
		return "", errDosen
	}

	var dosenLoads int
	dosenLoads = dosen.Load - matkul.SKS
	var dosenNew models.Dosen

	dosenNew.Nama = dosen.Nama
	dosenNew.KodeDosen = dosen.KodeDosen
	dosenNew.Preferensi = dosen.Preferensi
	dosenNew.RumpunID = dosen.RumpunID
	dosenNew.Load = dosenLoads

	dosenUpdateLoad, errDosenUpdate := p.dosenRepo.UpdateDosen(int(dosen.ID), dosenNew)
	response, err := p.perkuliahanRepo.DeletePerkuliahan(PerkuliahanID)
	println(dosenUpdateLoad, errDosenUpdate)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *perkuliahanService) UpdatePerkuliahan(PerkuliahanID int, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	Perkuliahan := models.Perkuliahan{
		Sesi:         PerkuliahanData.Sesi,
		Kelas:        PerkuliahanData.Kelas,
		Ruangan:      PerkuliahanData.Ruangan,
		MataKuliahId: PerkuliahanData.MataKuliahId,
		DosenId:      PerkuliahanData.DosenId,
	}

	PerkuliahanUpdated, err := p.perkuliahanRepo.UpdatePerkuliahan(PerkuliahanID, Perkuliahan)

	if err != nil {
		return nil, err
	}

	return PerkuliahanUpdated, nil
}

func (p *perkuliahanService) GetPerkuliahanById(PerkuliahanID int) (*models.Perkuliahan, errs.Errs) {
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
		matkul, errMatkul := p.matkulRepo.GetMatkulById(perkuliahan.MataKuliahId)
		if errMatkul != nil {
			return nil, errMatkul
		}
		dosen, errDosen := p.dosenRepo.GetDosenById(perkuliahan.DosenId)
		if errDosen != nil {
			return nil, errDosen
		}
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkul.RumpunID)
		if errRumpun != nil {
			return nil, errRumpun
		}

		JadwalResponse := dto.JadwalNewResponse{
			KodeDosen:      dosen.KodeDosen,
			KodeMataKuliah: matkul.KodeMK,
			Kelas:          perkuliahan.Kelas,
			Ruangan:        perkuliahan.Ruangan,
			Sesi:           perkuliahan.Sesi,
			Preferensi:     dto.Preferensi(dosen.Preferensi),
			Tipe:           matkul.Tipe,
			Rumpun:         rumpun.KodeRMK,
		}
		JadwalNewResponse = append(JadwalNewResponse, JadwalResponse)
	}

	if err != nil {
		return nil, err
	}

	return &JadwalNewResponse, nil
}

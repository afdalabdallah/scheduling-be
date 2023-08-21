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
	PerkuliahanCreated, err := p.perkuliahanRepo.CreatePerkuliahan(Perkuliahan)
	if err != nil {
		return nil, err
	}

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
	response, err := p.perkuliahanRepo.DeletePerkuliahan(PerkuliahanID)
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

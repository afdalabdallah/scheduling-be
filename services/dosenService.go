package services

import (
	"github.com/afdalabdallah/backend-web/dto"
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/dosen_repository"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
)

type dosenService struct {
	dosenRepo  dosen_repository.DosenRepository
	rumpunRepo rumpun_repository.RumpunRepository
}

type DosenService interface {
	CreateDosen(dosenData []models.Dosen) (*[]models.Dosen, errs.Errs)
	GetAllDosen() (*[]dto.DosenResponse, errs.Errs)
	DeleteDosen(dosenID int) (string, errs.Errs)
	UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs)
	GetDosenById(dosenID int) (*dto.DosenResponse, errs.Errs)
}

func NewDosenService(dosenRepo dosen_repository.DosenRepository, rumpunRepo rumpun_repository.RumpunRepository) DosenService {
	return &dosenService{
		dosenRepo:  dosenRepo,
		rumpunRepo: rumpunRepo,
	}
}

func (p *dosenService) CreateDosen(dosenData []models.Dosen) (*[]models.Dosen, errs.Errs) {
	var dosenCreateResponse []models.Dosen
	for _, data := range dosenData {
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(data.RumpunID)
		if errRumpun != nil || rumpun.ID == 0 {
			return nil, errRumpun
		}

		dosen := models.Dosen{
			Nama:       data.Nama,
			KodeDosen:  data.KodeDosen,
			Preferensi: data.Preferensi,
			RumpunID:   data.RumpunID,
			Load:       0,
		}
		dosenCreated, err := p.dosenRepo.CreateDosen(dosen)
		if err != nil {
			return nil, err
		}
		dosenCreateResponse = append(dosenCreateResponse, *dosenCreated)
	}

	return &dosenCreateResponse, nil
}

func (p *dosenService) GetAllDosen() (*[]dto.DosenResponse, errs.Errs) {
	dosens, err := p.dosenRepo.GetAllDosen()

	if err != nil {
		return nil, err
	}

	var dosenResponseArr []dto.DosenResponse
	for _, dosen := range dosens {
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(dosen.RumpunID)
		if errRumpun != nil {
			return nil, errRumpun
		}
		dosenResponse := dto.DosenResponse{
			ID:         int(dosen.ID),
			Nama:       dosen.Nama,
			KodeDosen:  dosen.KodeDosen,
			Preferensi: dto.Preferensi(dosen.Preferensi),
			Rumpun:     rumpun.KodeRMK,
			Load:       dosen.Load,
		}
		dosenResponseArr = append(dosenResponseArr, dosenResponse)
	}

	return &dosenResponseArr, nil
}

func (p *dosenService) DeleteDosen(dosenID int) (string, errs.Errs) {
	response, err := p.dosenRepo.DeleteDosen(dosenID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *dosenService) UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs) {
	rumpun, errRumpun := p.rumpunRepo.GetRumpunById(dosenData.RumpunID)
	if errRumpun != nil || rumpun.ID == 0 {
		return nil, errRumpun
	}

	dosen := models.Dosen{
		Nama:       dosenData.Nama,
		KodeDosen:  dosenData.KodeDosen,
		Preferensi: dosenData.Preferensi,
		RumpunID:   dosenData.RumpunID,
	}

	dosenUpdated, err := p.dosenRepo.UpdateDosen(dosenID, dosen)

	if err != nil {
		return nil, err
	}

	return dosenUpdated, nil
}

func (p *dosenService) GetDosenById(dosenID int) (*dto.DosenResponse, errs.Errs) {
	dosenData, err := p.dosenRepo.GetDosenById(dosenID)
	rumpun, errRumpun := p.rumpunRepo.GetRumpunById(dosenData.RumpunID)
	if errRumpun != nil || rumpun.ID == 0 {
		return nil, errRumpun
	}
	if err != nil {
		return nil, err
	}
	var dosenResponse dto.DosenResponse
	dosenResponse = dto.DosenResponse{
		ID:         int(dosenData.ID),
		Nama:       dosenData.Nama,
		KodeDosen:  dosenData.KodeDosen,
		Preferensi: dto.Preferensi(dosenData.Preferensi),
		Rumpun:     rumpun.KodeRMK,
	}

	return &dosenResponse, nil
}

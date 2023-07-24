package services

import (
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
	CreateDosen(dosenData models.Dosen) (*models.Dosen, errs.Errs)
	GetAllDosen() (*[]models.Dosen, errs.Errs)
	DeleteDosen(dosenID int) (string, errs.Errs)
	UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs)
	GetDosenById(dosenID int) (*models.Dosen, errs.Errs)
}

func NewDosenService(dosenRepo dosen_repository.DosenRepository, rumpunRepo rumpun_repository.RumpunRepository) DosenService {
	return &dosenService{
		dosenRepo:  dosenRepo,
		rumpunRepo: rumpunRepo,
	}
}

func (p *dosenService) CreateDosen(dosenData models.Dosen) (*models.Dosen, errs.Errs) {
	dosen := models.Dosen{
		Nama:       dosenData.Nama,
		KodeDosen:  dosenData.KodeDosen,
		Preferensi: dosenData.Preferensi,
		RumpunID:   dosenData.RumpunID,
	}
	dosenCreated, err := p.dosenRepo.CreateDosen(dosen)
	if err != nil {
		return nil, err
	}

	return dosenCreated, nil
}

func (p *dosenService) GetAllDosen() (*[]models.Dosen, errs.Errs) {
	dosens, err := p.dosenRepo.GetAllDosen()

	if err != nil {
		return nil, err
	}

	return &dosens, nil
}

func (p *dosenService) DeleteDosen(dosenID int) (string, errs.Errs) {
	response, err := p.dosenRepo.DeleteDosen(dosenID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *dosenService) UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs) {
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

func (p *dosenService) GetDosenById(dosenID int) (*models.Dosen, errs.Errs) {
	dosenData, err := p.dosenRepo.GetDosenById(dosenID)

	if err != nil {
		return nil, err
	}

	return dosenData, nil
}

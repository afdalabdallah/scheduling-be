package services

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/matkul_repository"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
)

type matkulService struct {
	matkulRepo matkul_repository.MatkulRepository
	rumpunRepo rumpun_repository.RumpunRepository
}

type MatkulService interface {
	CreateMatkul(matkulData models.Matkul) (*models.Matkul, errs.Errs)
	GetAllMatkul() (*[]models.Matkul, errs.Errs)
	DeleteMatkul(matkulID int) (string, errs.Errs)
	UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs)
	GetMatkulById(matkulID int) (*models.Matkul, errs.Errs)
}

func NewMatkulService(matkulRepo matkul_repository.MatkulRepository, rumpunRepo rumpun_repository.RumpunRepository) MatkulService {
	return &matkulService{
		matkulRepo: matkulRepo,
		rumpunRepo: rumpunRepo,
	}
}

func (p *matkulService) CreateMatkul(matkulData models.Matkul) (*models.Matkul, errs.Errs) {
	matkul := models.Matkul{
		Nama:     matkulData.Nama,
		KodeMK:   matkulData.KodeMK,
		Tipe:     matkulData.Tipe,
		Semester: matkulData.Semester,
		RumpunID: matkulData.RumpunID,
	}
	matkulCreated, err := p.matkulRepo.CreateMatkul(matkul)
	if err != nil {
		return nil, err
	}

	return matkulCreated, nil
}

func (p *matkulService) GetAllMatkul() (*[]models.Matkul, errs.Errs) {
	matkuls, err := p.matkulRepo.GetAllMatkul()

	if err != nil {
		return nil, err
	}

	return &matkuls, nil
}

func (p *matkulService) DeleteMatkul(matkulID int) (string, errs.Errs) {
	response, err := p.matkulRepo.DeleteMatkul(matkulID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *matkulService) UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs) {
	matkul := models.Matkul{
		Nama:     matkulData.Nama,
		KodeMK:   matkulData.KodeMK,
		Tipe:     matkulData.Tipe,
		Semester: matkulData.Semester,
		RumpunID: matkulData.RumpunID,
	}

	matkulUpdated, err := p.matkulRepo.UpdateMatkul(matkulID, matkul)

	if err != nil {
		return nil, err
	}

	return matkulUpdated, nil
}

func (p *matkulService) GetMatkulById(matkulID int) (*models.Matkul, errs.Errs) {
	matkulData, err := p.matkulRepo.GetMatkulById(matkulID)

	if err != nil {
		return nil, err
	}

	return matkulData, nil
}

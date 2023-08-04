package services

import (
	"github.com/afdalabdallah/backend-web/dto"
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
	GetAllMatkul() (*dto.MatkulResponse, errs.Errs)
	DeleteMatkul(matkulID int) (string, errs.Errs)
	UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs)
	GetMatkulById(matkulID int) (*dto.MatkulResponse, errs.Errs)
}

func NewMatkulService(matkulRepo matkul_repository.MatkulRepository, rumpunRepo rumpun_repository.RumpunRepository) MatkulService {
	return &matkulService{
		matkulRepo: matkulRepo,
		rumpunRepo: rumpunRepo,
	}
}

func (p *matkulService) CreateMatkul(matkulData models.Matkul) (*models.Matkul, errs.Errs) {
	rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkulData.RumpunID)
	if errRumpun != nil || rumpun.ID == 0 {
		return nil, errRumpun
	}

	matkul := models.Matkul{
		Nama:     matkulData.Nama,
		KodeMK:   matkulData.KodeMK,
		Tipe:     matkulData.Tipe,
		Semester: matkulData.Semester,
		RumpunID: matkulData.RumpunID,
		SKS:      matkulData.SKS,
	}
	matkulCreated, err := p.matkulRepo.CreateMatkul(matkul)
	if err != nil {
		return nil, err
	}

	return matkulCreated, nil
}

func (p *matkulService) GetAllMatkul() (*dto.MatkulResponse, errs.Errs) {
	matkuls, err := p.matkulRepo.GetAllMatkul()
	if err != nil {
		return nil, err
	}

	var matkulRespons dto.MatkulResponse
	for _, matkul := range matkuls {
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkul.RumpunID)
		if errRumpun != nil {
			return nil, errRumpun
		}

		matkulRespons = dto.MatkulResponse{
			Nama:     matkul.Nama,
			KodeMK:   matkul.KodeMK,
			Tipe:     matkul.Tipe,
			Semester: matkul.Semester,
			Rumpun:   rumpun.KodeRMK,
			SKS:      matkul.SKS,
		}
	}

	return &matkulRespons, nil
}

func (p *matkulService) DeleteMatkul(matkulID int) (string, errs.Errs) {
	response, err := p.matkulRepo.DeleteMatkul(matkulID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *matkulService) UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs) {
	rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkulData.RumpunID)
	if errRumpun != nil || rumpun.ID == 0 {
		return nil, errRumpun
	}

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

func (p *matkulService) GetMatkulById(matkulID int) (*dto.MatkulResponse, errs.Errs) {
	matkulData, err := p.matkulRepo.GetMatkulById(matkulID)

	if err != nil {
		return nil, err
	}
	rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkulData.RumpunID)
	if errRumpun != nil {
		return nil, errRumpun
	}

	matkulRespons := dto.MatkulResponse{
		Nama:     matkulData.Nama,
		KodeMK:   matkulData.KodeMK,
		Tipe:     matkulData.Tipe,
		Semester: matkulData.Semester,
		Rumpun:   rumpun.Nama,
		SKS:      matkulData.SKS,
	}

	return &matkulRespons, nil
}

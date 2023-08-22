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
	CreateMatkul(matkulData []models.Matkul) (*[]models.Matkul, errs.Errs)
	GetAllMatkul() (*[]dto.MatkulResponse, errs.Errs)
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

func (p *matkulService) CreateMatkul(matkulData []models.Matkul) (*[]models.Matkul, errs.Errs) {
	var matkulCreateRespons []models.Matkul
	for _, data := range matkulData {
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(data.RumpunID)
		if errRumpun != nil || rumpun.ID == 0 {
			return nil, errRumpun
		}

		matkul := models.Matkul{
			Nama:     data.Nama,
			KodeMK:   data.KodeMK,
			Tipe:     data.Tipe,
			Semester: data.Semester,
			RumpunID: data.RumpunID,
			SKS:      data.SKS,
		}
		matkulCreated, err := p.matkulRepo.CreateMatkul(matkul)
		if err != nil {
			return nil, err
		}
		matkulCreateRespons = append(matkulCreateRespons, *matkulCreated)

	}

	return &matkulCreateRespons, nil
}

func (p *matkulService) GetAllMatkul() (*[]dto.MatkulResponse, errs.Errs) {
	matkuls, err := p.matkulRepo.GetAllMatkul()
	if err != nil {
		return nil, err
	}

	var matkulRespons []dto.MatkulResponse
	for _, matkul := range matkuls {
		rumpun, errRumpun := p.rumpunRepo.GetRumpunById(matkul.RumpunID)
		if errRumpun != nil {
			return nil, errRumpun
		}

		matkulResponsData := dto.MatkulResponse{
			ID:       int(matkul.ID),
			Nama:     matkul.Nama,
			KodeMK:   matkul.KodeMK,
			Tipe:     matkul.Tipe,
			Semester: matkul.Semester,
			Rumpun:   rumpun.KodeRMK,
			SKS:      matkul.SKS,
		}
		matkulRespons = append(matkulRespons, matkulResponsData)
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
		ID:       int(matkulData.ID),
		Nama:     matkulData.Nama,
		KodeMK:   matkulData.KodeMK,
		Tipe:     matkulData.Tipe,
		Semester: matkulData.Semester,
		Rumpun:   rumpun.Nama,
		SKS:      matkulData.SKS,
	}

	return &matkulRespons, nil
}

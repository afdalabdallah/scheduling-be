package services

import (
	"github.com/afdalabdallah/backend-web/dto"
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/ruangan_repository"
)

type ruanganService struct {
	ruanganRepo ruangan_repository.RuanganRepository
}

type RuanganService interface {
	CreateRuangan(ruanganData []models.Ruangan) (*[]models.Ruangan, errs.Errs)
	GetAllRuangan() (*[]dto.RuanganResponse, errs.Errs)
	DeleteRuangan(ruanganID uint) (string, errs.Errs)
	UpdateRuangan(ruanganID uint, ruanganData models.Ruangan) (*models.Ruangan, errs.Errs)
	GetRuanganById(ruanganID uint) (*models.Ruangan, errs.Errs)
}

func NewRuanganService(ruanganRepo ruangan_repository.RuanganRepository) RuanganService {
	return &ruanganService{
		ruanganRepo: ruanganRepo,
	}
}

func (p *ruanganService) CreateRuangan(ruanganData []models.Ruangan) (*[]models.Ruangan, errs.Errs) {
	var ruanganCreateResponse []models.Ruangan
	for _, data := range ruanganData {
		ruangan := models.Ruangan{
			Nomor:     data.Nomor,
			Kapasitas: data.Kapasitas,
			Deskripsi: data.Deskripsi,
		}
		ruanganCreated, err := p.ruanganRepo.CreateRuangan(ruangan)
		if err != nil {
			return nil, err
		}
		ruanganCreateResponse = append(ruanganCreateResponse, *ruanganCreated)

	}

	return &ruanganCreateResponse, nil
}

func (p *ruanganService) GetAllRuangan() (*[]dto.RuanganResponse, errs.Errs) {
	ruangans, err := p.ruanganRepo.GetAllRuangan()

	var ruanganResponse []dto.RuanganResponse
	for _, ruangan := range ruangans {
		ruanganRes := dto.RuanganResponse{
			ID:        int(ruangan.ID),
			Nomor:     ruangan.Nomor,
			Kapasitas: ruangan.Kapasitas,
			Deskripsi: ruangan.Deskripsi,
		}
		ruanganResponse = append(ruanganResponse, ruanganRes)
	}

	if err != nil {
		return nil, err
	}

	return &ruanganResponse, nil
}

func (p *ruanganService) DeleteRuangan(ruanganID uint) (string, errs.Errs) {
	response, err := p.ruanganRepo.DeleteRuangan(ruanganID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *ruanganService) UpdateRuangan(ruanganID uint, ruanganData models.Ruangan) (*models.Ruangan, errs.Errs) {
	ruangan := models.Ruangan{
		Nomor:     ruanganData.Nomor,
		Kapasitas: ruanganData.Kapasitas,
		Deskripsi: ruanganData.Deskripsi,
	}

	ruanganUpdated, err := p.ruanganRepo.UpdateRuangan(ruanganID, ruangan)

	if err != nil {
		return nil, err
	}

	return ruanganUpdated, nil
}

func (p *ruanganService) GetRuanganById(ruanganID uint) (*models.Ruangan, errs.Errs) {
	ruanganData, err := p.ruanganRepo.GetRuanganById(ruanganID)

	if err != nil {
		return nil, err
	}

	return ruanganData, nil
}

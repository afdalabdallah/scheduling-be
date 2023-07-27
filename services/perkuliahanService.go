package services

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/perkuliahan_repository"
)

type perkuliahanService struct {
	perkuliahanRepo perkuliahan_repository.PerkuliahanRepository
}

type PerkuliahanService interface {
	CreatePerkuliahan(PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetAllPerkuliahan() (*[]models.Perkuliahan, errs.Errs)
	DeletePerkuliahan(PerkuliahanID int) (string, errs.Errs)
	UpdatePerkuliahan(PerkuliahanID int, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanById(PerkuliahanID int) (*models.Perkuliahan, errs.Errs)
}

func NewPerkuliahanService(perkuliahanRepo perkuliahan_repository.PerkuliahanRepository) PerkuliahanService {
	return &perkuliahanService{
		perkuliahanRepo: perkuliahanRepo,
	}
}

func (p *perkuliahanService) CreatePerkuliahan(PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	Perkuliahan := models.Perkuliahan{
		Sesi:     PerkuliahanData.Sesi,
		Kelas:   PerkuliahanData.Kelas,
		Ruangan:     PerkuliahanData.Ruangan,
		MataKuliahId: PerkuliahanData.MataKuliahId,
	}
	PerkuliahanCreated, err := p.perkuliahanRepo.CreatePerkuliahan(Perkuliahan)
	if err != nil {
		return nil, err
	}

	return PerkuliahanCreated, nil
}

func (p *perkuliahanService) GetAllPerkuliahan() (*[]models.Perkuliahan, errs.Errs) {
	Perkuliahans, err := p.perkuliahanRepo.GetAllPerkuliahan()

	if err != nil {
		return nil, err
	}

	return &Perkuliahans, nil
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
		Sesi:     PerkuliahanData.Sesi,
		Kelas:   PerkuliahanData.Kelas,
		Ruangan:     PerkuliahanData.Ruangan,
		MataKuliahId: PerkuliahanData.MataKuliahId,
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

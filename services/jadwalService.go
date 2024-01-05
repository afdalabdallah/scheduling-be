package services

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/jadwal_repository"
)

type jadwalService struct {
	jadwalRepo jadwal_repository.JadwalRepository
}

type JadwalService interface {
	CreateJadwal(jadwalData models.Jadwal) (*models.Jadwal, errs.Errs)
	GetAllJadwal() (*[]models.Jadwal, errs.Errs)
	DeleteJadwal(jadwalID uint) (string, errs.Errs)
	UpdateJadwal(jadwalID uint, jadwalData models.Jadwal) (*models.Jadwal, errs.Errs)
	GetJadwalById(jadwalID uint) (*models.Jadwal, errs.Errs)
	// GetLatestJadwal() (*models.Jadwal, errs.Errs)
}

func NewJadwalService(jadwalRepo jadwal_repository.JadwalRepository) JadwalService {
	return &jadwalService{
		jadwalRepo: jadwalRepo,
	}
}

func (p *jadwalService) CreateJadwal(jadwalData models.Jadwal) (*models.Jadwal, errs.Errs) {
	var jadwalCreateResponse models.Jadwal
	jadwal := models.Jadwal{
		Data:               jadwalData.Data,
		Fitness:            jadwalData.Fitness,
		ViolatedConstraint: jadwalData.ViolatedConstraint,
		Skpb:               jadwalData.Skpb,
		UnwantedSesi:	jadwalData.UnwantedSesi,
		ListRuangan: jadwalData.ListRuangan,
	}
	jadwalCreated, err := p.jadwalRepo.CreateJadwal(jadwal)
	if err != nil {
		return nil, err
	}
	jadwalCreateResponse = *jadwalCreated

	return &jadwalCreateResponse, nil
}

func (p *jadwalService) GetAllJadwal() (*[]models.Jadwal, errs.Errs) {
	jadwals, err := p.jadwalRepo.GetAllJadwal()

	if err != nil {
		return nil, err
	}

	return &jadwals, nil
}

func (p *jadwalService) DeleteJadwal(jadwalID uint) (string, errs.Errs) {
	response, err := p.jadwalRepo.DeleteJadwal(jadwalID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *jadwalService) UpdateJadwal(jadwalID uint, jadwalData models.Jadwal) (*models.Jadwal, errs.Errs) {

	jadwal := models.Jadwal{
		Data:               jadwalData.Data,
		Fitness:            jadwalData.Fitness,
		ViolatedConstraint: jadwalData.ViolatedConstraint,
		Skpb:               jadwalData.Skpb,
		UnwantedSesi:	jadwalData.UnwantedSesi,
		ListRuangan: jadwalData.ListRuangan,
	}

	jadwalUpdated, err := p.jadwalRepo.UpdateJadwal(jadwalID, jadwal)

	if err != nil {
		return nil, err
	}

	return jadwalUpdated, nil
}

func (p *jadwalService) GetJadwalById(jadwalID uint) (*models.Jadwal, errs.Errs) {
	jadwalData, err := p.jadwalRepo.GetJadwalById(jadwalID)

	if err != nil {
		return nil, err
	}
	var jadwalResponse models.Jadwal
	jadwalResponse = models.Jadwal{
		Data:               jadwalData.Data,
		Fitness:            jadwalData.Fitness,
		ViolatedConstraint: jadwalData.ViolatedConstraint,
		Skpb:               jadwalData.Skpb,
		UnwantedSesi:	jadwalData.UnwantedSesi,
		ListRuangan: jadwalData.ListRuangan,
		
	}

	return &jadwalResponse, nil
}

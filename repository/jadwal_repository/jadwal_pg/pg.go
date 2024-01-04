package jadwal_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/jadwal_repository"
	"gorm.io/gorm"
)

type jadwalRepository struct {
	db *gorm.DB
}

func NewPGJadwalRepository(db *gorm.DB) jadwal_repository.JadwalRepository {
	return &jadwalRepository{
		db: db,
	}
}

func (p *jadwalRepository) CreateJadwal(jadwal models.Jadwal) (*models.Jadwal, errs.Errs) {
	res := p.db.Create(&jadwal)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &jadwal, nil
}

func (p *jadwalRepository) GetAllJadwal() ([]models.Jadwal, errs.Errs) {
	var jadwal []models.Jadwal

	result := p.db.Find(&jadwal)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return jadwal, nil
}

func (p *jadwalRepository) DeleteJadwal(jadwalID uint) (string, errs.Errs) {
	result := p.db.Delete(&models.Jadwal{}, jadwalID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Jadwal has been successfully deleted", nil
}

func (p *jadwalRepository) UpdateJadwal(jadwalID uint, jadwalData models.Jadwal) (*models.Jadwal, errs.Errs) {
	var jadwalUpdate models.Jadwal

	// Get data by id
	result := p.db.First(&jadwalUpdate, jadwalID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	jadwalUpdate.Data = jadwalData.Data
	jadwalUpdate.Fitness = jadwalData.Fitness
	jadwalUpdate.ViolatedConstraint = jadwalData.ViolatedConstraint
	jadwalUpdate.Skpb = jadwalData.Skpb

	result = p.db.Save(&jadwalUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &jadwalUpdate, nil
}

func (p *jadwalRepository) GetJadwalById(jadwalID uint) (*models.Jadwal, errs.Errs) {
	var jadwalData models.Jadwal

	result := p.db.Preload("Rumpun").First(&jadwalData, jadwalID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &jadwalData, nil
}

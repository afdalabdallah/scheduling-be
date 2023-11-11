package ruangan_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/ruangan_repository"
	"gorm.io/gorm"
)

type ruanganRepository struct {
	db *gorm.DB
}

func NewPGRuanganRepository(db *gorm.DB) ruangan_repository.RuanganRepository {
	return &ruanganRepository{
		db: db,
	}
}

func (p *ruanganRepository) CreateRuangan(ruangan models.Ruangan) (*models.Ruangan, errs.Errs) {
	res := p.db.Create(&ruangan)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &ruangan, nil
}

func (p *ruanganRepository) GetAllRuangan() ([]models.Ruangan, errs.Errs) {
	var ruangan []models.Ruangan

	result := p.db.Find(&ruangan)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return ruangan, nil
}

func (p *ruanganRepository) DeleteRuangan(ruanganID uint) (string, errs.Errs) {
	result := p.db.Unscoped().Delete(&models.Ruangan{}, ruanganID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Ruangan has been successfully deleted", nil
}

func (p *ruanganRepository) UpdateRuangan(ruanganID uint, ruanganData models.Ruangan) (*models.Ruangan, errs.Errs) {
	var ruanganUpdate models.Ruangan

	// Get data by id
	result := p.db.First(&ruanganUpdate, ruanganID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	ruanganUpdate.Nomor = ruanganData.Nomor
	ruanganUpdate.Kapasitas = ruanganData.Kapasitas
	ruanganUpdate.Deskripsi = ruanganData.Deskripsi

	result = p.db.Save(&ruanganUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &ruanganUpdate, nil
}

func (p *ruanganRepository) GetRuanganById(ruanganID uint) (*models.Ruangan, errs.Errs) {
	var ruanganData models.Ruangan

	result := p.db.First(&ruanganData, ruanganID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &ruanganData, nil
}

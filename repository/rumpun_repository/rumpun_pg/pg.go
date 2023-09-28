package rumpun_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
	"gorm.io/gorm"
)

type rumpunRepository struct {
	db *gorm.DB
}

func NewPGRumpunRepository(db *gorm.DB) rumpun_repository.RumpunRepository {
	return &rumpunRepository{
		db: db,
	}
}

// Create RMK
func (p *rumpunRepository) CreateRumpun(rumpun models.Rumpun) (*models.Rumpun, errs.Errs) {
	res := p.db.Create(&rumpun)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &rumpun, nil
}

// Get All RMK
func (p *rumpunRepository) GetAllRumpun() ([]models.Rumpun, errs.Errs) {
	var rumpun []models.Rumpun

	result := p.db.Find(&rumpun)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return rumpun, nil
}

// Delete RMK
func (p *rumpunRepository) DeleteRMK(rumpunID uint) (string, errs.Errs) {
	result := p.db.Unscoped().Delete(&models.Rumpun{}, rumpunID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Product has been successfully deleted", nil
}

// Update RMK
func (p *rumpunRepository) UpdateRMK(rumpunID uint, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs) {
	var rumpunUpdate models.Rumpun

	// Get data by id
	result := p.db.First(&rumpunUpdate, rumpunID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	rumpunUpdate.Nama = rumpunData.Nama
	rumpunUpdate.KodeRMK = rumpunData.KodeRMK

	result = p.db.Save(&rumpunUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &rumpunUpdate, nil
}

func (p *rumpunRepository) GetRumpunById(rumpunID uint) (*models.Rumpun, errs.Errs) {
	var rumpunData models.Rumpun

	result := p.db.First(&rumpunData, rumpunID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &rumpunData, nil
}

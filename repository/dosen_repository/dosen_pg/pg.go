package dosen_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/dosen_repository"
	"gorm.io/gorm"
)

type dosenRepository struct {
	db *gorm.DB
}

func NewPGDosenRepository(db *gorm.DB) dosen_repository.DosenRepository {
	return &dosenRepository{
		db: db,
	}
}

func (p *dosenRepository) CreateDosen(dosen models.Dosen) (*models.Dosen, errs.Errs) {
	res := p.db.Create(&dosen)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &dosen, nil
}

func (p *dosenRepository) GetAllDosen() ([]models.Dosen, errs.Errs) {
	var dosen []models.Dosen

	result := p.db.Find(&dosen)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return dosen, nil
}

func (p *dosenRepository) DeleteDosen(dosenID int) (string, errs.Errs) {
	result := p.db.Delete(&models.Dosen{}, dosenID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Mata Kuliah has been successfully deleted", nil
}

func (p *dosenRepository) UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs) {
	var dosenUpdate models.Dosen

	// Get data by id
	result := p.db.First(&dosenUpdate, dosenID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	dosenUpdate.Nama = dosenData.Nama
	dosenUpdate.KodeDosen = dosenData.KodeDosen
	dosenUpdate.Preferensi = dosenData.Preferensi
	dosenUpdate.RumpunID = dosenData.RumpunID
	dosenUpdate.Load = dosenData.Load

	result = p.db.Save(&dosenUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &dosenUpdate, nil
}

func (p *dosenRepository) GetDosenById(dosenID int) (*models.Dosen, errs.Errs) {
	var dosenData models.Dosen

	result := p.db.First(&dosenData, dosenID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &dosenData, nil
}

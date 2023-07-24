package matkul_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/matkul_repository"
	"gorm.io/gorm"
)

type matkulRepository struct {
	db *gorm.DB
}

func NewPGMatkulRepository(db *gorm.DB) matkul_repository.MatkulRepository {
	return &matkulRepository{
		db: db,
	}
}

func (p *matkulRepository) CreateMatkul(matkul models.Matkul) (*models.Matkul, errs.Errs) {
	res := p.db.Create(&matkul)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &matkul, nil
}

func (p *matkulRepository) GetAllMatkul() ([]models.Matkul, errs.Errs) {
	var matkul []models.Matkul

	result := p.db.Find(&matkul)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return matkul, nil
}

func (p *matkulRepository) DeleteMatkul(matkulID int) (string, errs.Errs) {
	result := p.db.Delete(&models.Matkul{}, matkulID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Mata Kuliah has been successfully deleted", nil
}

func (p *matkulRepository) UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs) {
	var matkulUpdate models.Matkul

	// Get data by id
	result := p.db.First(&matkulUpdate, matkulID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	matkulUpdate.Nama = matkulData.Nama
	matkulUpdate.KodeMK = matkulData.KodeMK
	matkulUpdate.Tipe = matkulData.Tipe
	matkulUpdate.Semester = matkulData.Semester
	matkulUpdate.RumpunID = matkulData.RumpunID

	result = p.db.Save(&matkulUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &matkulUpdate, nil
}

func (p *matkulRepository) GetMatkulById(matkulID int) (*models.Matkul, errs.Errs) {
	var matkulData models.Matkul

	result := p.db.First(&matkulData, matkulID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &matkulData, nil
}

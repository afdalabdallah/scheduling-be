package perkuliahan_pg

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/perkuliahan_repository"
	"gorm.io/gorm"
)

type perkuliahanRepository struct {
	db *gorm.DB
}

func NewPGPerkuliahanRepository(db *gorm.DB) perkuliahan_repository.PerkuliahanRepository {
	return &perkuliahanRepository{
		db: db,
	}
}

func (p *perkuliahanRepository) CreatePerkuliahan(Perkuliahan models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	res := p.db.Create(&Perkuliahan)
	err := res.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &Perkuliahan, nil
}

func (p *perkuliahanRepository) GetAllPerkuliahan() ([]models.Perkuliahan, errs.Errs) {
	var Perkuliahan []models.Perkuliahan
	result := p.db.Preload("Matkul").Preload("Rumpun").Preload("Dosen").Find(&Perkuliahan)
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return Perkuliahan, nil
}

func (p *perkuliahanRepository) DeletePerkuliahan(PerkuliahanID uint) (string, errs.Errs) {
	result := p.db.Delete(&models.Perkuliahan{}, PerkuliahanID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Mata Kuliah has been successfully deleted", nil
}

func (p *perkuliahanRepository) UpdatePerkuliahan(PerkuliahanID uint, PerkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs) {
	var PerkuliahanUpdate models.Perkuliahan

	// Get data by id
	result := p.db.First(&PerkuliahanUpdate, PerkuliahanID)

	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	PerkuliahanUpdate.Kelas = PerkuliahanData.Kelas
	PerkuliahanUpdate.Sesi = PerkuliahanData.Sesi
	PerkuliahanUpdate.Ruangan = PerkuliahanData.Ruangan
	PerkuliahanUpdate.MatkulID = PerkuliahanData.MatkulID
	PerkuliahanUpdate.RumpunID = PerkuliahanData.RumpunID

	result = p.db.Save(&PerkuliahanUpdate)

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &PerkuliahanUpdate, nil
}

func (p *perkuliahanRepository) GetPerkuliahanById(PerkuliahanID uint) (*models.Perkuliahan, errs.Errs) {
	var PerkuliahanData models.Perkuliahan

	result := p.db.Preload("Matkul").Preload("Rumpun").Preload("Dosen").Find(&PerkuliahanData)
	err := result.Error
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &PerkuliahanData, nil
}

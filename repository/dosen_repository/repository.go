package dosen_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type DosenRepository interface {
	CreateDosen(dosen models.Dosen) (*models.Dosen, errs.Errs)
	GetAllDosen() ([]models.Dosen, errs.Errs)
	DeleteDosen(dosenID uint) (string, errs.Errs)
	UpdateDosen(dosenID uint, dosenData models.Dosen) (*models.Dosen, errs.Errs)
	GetDosenById(dosenID uint) (*models.Dosen, errs.Errs)
}

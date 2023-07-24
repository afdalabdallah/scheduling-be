package dosen_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type DosenRepository interface {
	CreateDosen(dosen models.Dosen) (*models.Dosen, errs.Errs)
	GetAllDosen() ([]models.Dosen, errs.Errs)
	DeleteDosen(dosenID int) (string, errs.Errs)
	UpdateDosen(dosenID int, dosenData models.Dosen) (*models.Dosen, errs.Errs)
	GetDosenById(dosenID int) (*models.Dosen, errs.Errs)
}

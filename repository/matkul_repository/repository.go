package matkul_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type MatkulRepository interface {
	CreateMatkul(matkul models.Matkul) (*models.Matkul, errs.Errs)
	GetAllMatkul() ([]models.Matkul, errs.Errs)
	DeleteMatkul(matkulID int) (string, errs.Errs)
	UpdateMatkul(matkulID int, matkulData models.Matkul) (*models.Matkul, errs.Errs)
	GetMatkulById(matkulID int) (*models.Matkul, errs.Errs)
}

package perkuliahan_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type PerkuliahanRepository interface {
	CreatePerkuliahan(perkuliahan models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetAllPerkuliahan() ([]models.Perkuliahan, errs.Errs)
	DeletePerkuliahan(perkuliahanID uint) (string, errs.Errs)
	UpdatePerkuliahan(perkuliahanID uint, perkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanById(perkuliahanID uint) (*models.Perkuliahan, errs.Errs)
}

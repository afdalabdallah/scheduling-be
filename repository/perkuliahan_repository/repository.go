package perkuliahan_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type PerkuliahanRepository interface {
	CreatePerkuliahan(perkuliahan models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetAllPerkuliahan() ([]models.Perkuliahan, errs.Errs)
	DeletePerkuliahan(perkuliahanID int) (string, errs.Errs)
	UpdatePerkuliahan(perkuliahanID int, perkuliahanData models.Perkuliahan) (*models.Perkuliahan, errs.Errs)
	GetPerkuliahanById(perkuliahanID int) (*models.Perkuliahan, errs.Errs)
}

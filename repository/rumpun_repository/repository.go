package rumpun_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type RumpunRepository interface {
	CreateRumpun(rumpun models.Rumpun) (*models.Rumpun, errs.Errs)
	GetAllRumpun() ([]models.Rumpun, errs.Errs)
	DeleteRMK(rumpunID int) (string, errs.Errs)
	UpdateRMK(rumpunID int, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetRumpunById(rumpunID int) (*models.Rumpun, errs.Errs)
}

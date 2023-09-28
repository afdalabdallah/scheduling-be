package rumpun_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type RumpunRepository interface {
	CreateRumpun(rumpun models.Rumpun) (*models.Rumpun, errs.Errs)
	GetAllRumpun() ([]models.Rumpun, errs.Errs)
	DeleteRMK(rumpunID uint) (string, errs.Errs)
	UpdateRMK(rumpunID uint, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetRumpunById(rumpunID uint) (*models.Rumpun, errs.Errs)
}

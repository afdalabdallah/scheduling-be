package ruangan_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type RuanganRepository interface {
	CreateRuangan(rumpun models.Ruangan) (*models.Ruangan, errs.Errs)
	GetAllRuangan() ([]models.Ruangan, errs.Errs)
	DeleteRuangan(rumpunID uint) (string, errs.Errs)
	UpdateRuangan(rumpunID uint, rumpunData models.Ruangan) (*models.Ruangan, errs.Errs)
	GetRuanganById(rumpunID uint) (*models.Ruangan, errs.Errs)
}

package jadwal_repository

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
)

type JadwalRepository interface {
	CreateJadwal(jadwal models.Jadwal) (*models.Jadwal, errs.Errs)
	GetAllJadwal() ([]models.Jadwal, errs.Errs)
	DeleteJadwal(jadwalID uint) (string, errs.Errs)
	UpdateJadwal(jadwalID uint, jadwalData models.Jadwal) (*models.Jadwal, errs.Errs)
	GetJadwalById(jadwalID uint) (*models.Jadwal, errs.Errs)
}

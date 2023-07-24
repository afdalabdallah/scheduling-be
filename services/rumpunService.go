package services

import (
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
)

type rumpunService struct {
	rumpunRepo rumpun_repository.RumpunRepository
}

type RumpunService interface {
	CreateRumpun(rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetAllRumpun() (*[]models.Rumpun, errs.Errs)
	DeleteRumpun(rumpunID int) (string, errs.Errs)
	UpdateRumpun(rumpunID int, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetRumpunById(rumpunID int) (*models.Rumpun, errs.Errs)
}

func NewRumpunService(rumpunRepo rumpun_repository.RumpunRepository) RumpunService {
	return &rumpunService{
		rumpunRepo: rumpunRepo,
	}
}

func (p *rumpunService) CreateRumpun(rumpunData models.Rumpun) (*models.Rumpun, errs.Errs) {
	rumpun := models.Rumpun{
		Nama:    rumpunData.Nama,
		KodeRMK: rumpunData.KodeRMK,
	}
	rumpunCreated, err := p.rumpunRepo.CreateRumpun(rumpun)
	if err != nil {
		return nil, err
	}

	return rumpunCreated, nil
}

func (p *rumpunService) GetAllRumpun() (*[]models.Rumpun, errs.Errs) {
	rumpuns, err := p.rumpunRepo.GetAllRumpun()

	if err != nil {
		return nil, err
	}

	return &rumpuns, nil
}

func (p *rumpunService) DeleteRumpun(rumpunID int) (string, errs.Errs) {
	response, err := p.rumpunRepo.DeleteRMK(rumpunID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *rumpunService) UpdateRumpun(rumpunID int, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs) {
	rumpun := models.Rumpun{
		Nama:    rumpunData.Nama,
		KodeRMK: rumpunData.KodeRMK,
	}

	rumpunUpdated, err := p.rumpunRepo.UpdateRMK(rumpunID, rumpun)

	if err != nil {
		return nil, err
	}

	return rumpunUpdated, nil
}

func (p *rumpunService) GetRumpunById(rumpunID int) (*models.Rumpun, errs.Errs) {
	rumpunData, err := p.rumpunRepo.GetRumpunById(rumpunID)

	if err != nil {
		return nil, err
	}

	return rumpunData, nil
}

package services

import (
	"github.com/afdalabdallah/backend-web/dto"
	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository"
)

type rumpunService struct {
	rumpunRepo rumpun_repository.RumpunRepository
}

type RumpunService interface {
	CreateRumpun(rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetAllRumpun() (*[]dto.RumpunResponse, errs.Errs)
	DeleteRumpun(rumpunID uint) (string, errs.Errs)
	UpdateRumpun(rumpunID uint, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs)
	GetRumpunById(rumpunID uint) (*models.Rumpun, errs.Errs)
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

func (p *rumpunService) GetAllRumpun() (*[]dto.RumpunResponse, errs.Errs) {
	rumpuns, err := p.rumpunRepo.GetAllRumpun()

	var rumpunResponse []dto.RumpunResponse
	for _, rumpun := range rumpuns {
		rumpunRes := dto.RumpunResponse{
			ID:      int(rumpun.ID),
			Nama:    rumpun.Nama,
			KodeRMK: rumpun.KodeRMK,
		}
		rumpunResponse = append(rumpunResponse, rumpunRes)
	}

	if err != nil {
		return nil, err
	}

	return &rumpunResponse, nil
}

func (p *rumpunService) DeleteRumpun(rumpunID uint) (string, errs.Errs) {
	response, err := p.rumpunRepo.DeleteRMK(rumpunID)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *rumpunService) UpdateRumpun(rumpunID uint, rumpunData models.Rumpun) (*models.Rumpun, errs.Errs) {
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

func (p *rumpunService) GetRumpunById(rumpunID uint) (*models.Rumpun, errs.Errs) {
	rumpunData, err := p.rumpunRepo.GetRumpunById(rumpunID)

	if err != nil {
		return nil, err
	}

	return rumpunData, nil
}

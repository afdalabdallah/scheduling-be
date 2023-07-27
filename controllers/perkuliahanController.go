package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type perkuliahanController struct {
	PerkuliahanService services.PerkuliahanService
}

type PerkuliahanController interface {
	CreatePerkuliahan(c *gin.Context)
	GetAllPerkuliahan(c *gin.Context)
	DeletePerkuliahan(c *gin.Context)
	UpdatePerkuliahan(c *gin.Context)
	GetPerkuliahanById(c *gin.Context)
}

func NewPerkuliahanController(PerkuliahanService services.PerkuliahanService) PerkuliahanController {
	return &perkuliahanController{
		PerkuliahanService: PerkuliahanService,
	}
}

func (p *perkuliahanController) CreatePerkuliahan(c *gin.Context) {
	var PerkuliahanData models.Perkuliahan
	if err := c.ShouldBindJSON(&PerkuliahanData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.PerkuliahanService.CreatePerkuliahan(PerkuliahanData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *perkuliahanController) GetAllPerkuliahan(c *gin.Context) {
	response, err := p.PerkuliahanService.GetAllPerkuliahan()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *perkuliahanController) DeletePerkuliahan(c *gin.Context) {
	PerkuliahanID, err := strconv.Atoi(c.Param("perkuliahanID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.PerkuliahanService.DeletePerkuliahan(PerkuliahanID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *perkuliahanController) UpdatePerkuliahan(c *gin.Context) {
	PerkuliahanID, err := strconv.Atoi(c.Param("perkuliahanID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var PerkuliahanData models.Perkuliahan
	if err := c.ShouldBindJSON(&PerkuliahanData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.PerkuliahanService.UpdatePerkuliahan(PerkuliahanID, PerkuliahanData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *perkuliahanController) GetPerkuliahanById(c *gin.Context) {
	PerkuliahanID, err := strconv.Atoi(c.Param("perkuliahanID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.PerkuliahanService.GetPerkuliahanById(PerkuliahanID)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(200, response)
}

package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type ruanganController struct {
	ruanganService services.RuanganService
}

type RuanganController interface {
	CreateRuangan(c *gin.Context)
	GetAllRuangan(c *gin.Context)
	DeleteRuangan(c *gin.Context)
	UpdateRuangan(c *gin.Context)
	GetRuanganById(c *gin.Context)
}

func NewRuanganController(ruanganService services.RuanganService) RuanganController {
	return &ruanganController{
		ruanganService: ruanganService,
	}
}

func (p *ruanganController) CreateRuangan(c *gin.Context) {
	var ruanganData []models.Ruangan
	if err := c.ShouldBindJSON(&ruanganData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.ruanganService.CreateRuangan(ruanganData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *ruanganController) GetAllRuangan(c *gin.Context) {
	response, err := p.ruanganService.GetAllRuangan()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *ruanganController) DeleteRuangan(c *gin.Context) {
	ruanganID, err := strconv.Atoi(c.Param("ruanganID"))
	// println(ruanganID)
	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.ruanganService.DeleteRuangan(uint(ruanganID))

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *ruanganController) UpdateRuangan(c *gin.Context) {
	ruanganID, err := strconv.Atoi(c.Param("ruanganID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var ruanganData models.Ruangan
	if err := c.ShouldBindJSON(&ruanganData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.ruanganService.UpdateRuangan(uint(ruanganID), ruanganData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *ruanganController) GetRuanganById(c *gin.Context) {
	ruanganID, err := strconv.Atoi(c.Param("ruanganID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.ruanganService.GetRuanganById(uint(ruanganID))

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(200, response)
}

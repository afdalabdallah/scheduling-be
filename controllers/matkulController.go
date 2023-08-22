package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type matkulController struct {
	matkulService services.MatkulService
}

type MatkulController interface {
	CreateMatkul(c *gin.Context)
	GetAllMatkul(c *gin.Context)
	DeleteMatkul(c *gin.Context)
	UpdateMatkul(c *gin.Context)
	GetMatkulById(c *gin.Context)
}

func NewMatkulController(matkulService services.MatkulService) MatkulController {
	return &matkulController{
		matkulService: matkulService,
	}
}

func (p *matkulController) CreateMatkul(c *gin.Context) {
	var matkulData []models.Matkul
	if err := c.ShouldBindJSON(&matkulData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.matkulService.CreateMatkul(matkulData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *matkulController) GetAllMatkul(c *gin.Context) {
	response, err := p.matkulService.GetAllMatkul()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *matkulController) DeleteMatkul(c *gin.Context) {
	matkulID, err := strconv.Atoi(c.Param("matkulID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.matkulService.DeleteMatkul(matkulID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *matkulController) UpdateMatkul(c *gin.Context) {
	matkulID, err := strconv.Atoi(c.Param("matkulID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var matkulData models.Matkul
	if err := c.ShouldBindJSON(&matkulData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.matkulService.UpdateMatkul(matkulID, matkulData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *matkulController) GetMatkulById(c *gin.Context) {
	matkulID, err := strconv.Atoi(c.Param("matkulID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.matkulService.GetMatkulById(matkulID)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(200, response)
}

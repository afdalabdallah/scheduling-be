package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type dosenController struct {
	dosenService services.DosenService
}

type DosenController interface {
	CreateDosen(c *gin.Context)
	GetAllDosen(c *gin.Context)
	DeleteDosen(c *gin.Context)
	UpdateDosen(c *gin.Context)
	GetDosenById(c *gin.Context)
}

func NewDosenController(dosenService services.DosenService) DosenController {
	return &dosenController{
		dosenService: dosenService,
	}
}

func (p *dosenController) CreateDosen(c *gin.Context) {
	var dosenData []models.Dosen
	if err := c.ShouldBindJSON(&dosenData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.dosenService.CreateDosen(dosenData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *dosenController) GetAllDosen(c *gin.Context) {
	response, err := p.dosenService.GetAllDosen()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *dosenController) DeleteDosen(c *gin.Context) {
	dosenID, err := strconv.Atoi(c.Param("dosenID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.dosenService.DeleteDosen(dosenID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *dosenController) UpdateDosen(c *gin.Context) {
	dosenID, err := strconv.Atoi(c.Param("dosenID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var dosenData models.Dosen
	if err := c.ShouldBindJSON(&dosenData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.dosenService.UpdateDosen(dosenID, dosenData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *dosenController) GetDosenById(c *gin.Context) {
	dosenID, err := strconv.Atoi(c.Param("dosenID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.dosenService.GetDosenById(dosenID)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		print("msuk sini harusnya")
		return
	}

	c.JSON(200, response)
}

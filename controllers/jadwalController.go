package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type jadwalController struct {
	jadwalService services.JadwalService
}

type JadwalController interface {
	CreateJadwal(c *gin.Context)
	GetAllJadwal(c *gin.Context)
	DeleteJadwal(c *gin.Context)
	UpdateJadwal(c *gin.Context)
	GetJadwalById(c *gin.Context)
}

func NewJadwalController(jadwalService services.JadwalService) JadwalController {
	return &jadwalController{
		jadwalService: jadwalService,
	}
}

func (p *jadwalController) CreateJadwal(c *gin.Context) {
	var jadwalData models.Jadwal
	if err := c.ShouldBindJSON(&jadwalData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.jadwalService.CreateJadwal(jadwalData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *jadwalController) GetAllJadwal(c *gin.Context) {
	response, err := p.jadwalService.GetAllJadwal()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *jadwalController) DeleteJadwal(c *gin.Context) {
	jadwalID, err := strconv.Atoi(c.Param("jadwalID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.jadwalService.DeleteJadwal(uint(jadwalID))

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *jadwalController) UpdateJadwal(c *gin.Context) {
	jadwalID, err := strconv.Atoi(c.Param("jadwalID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var jadwalData models.Jadwal
	if err := c.ShouldBindJSON(&jadwalData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.jadwalService.UpdateJadwal(uint(jadwalID), jadwalData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *jadwalController) GetJadwalById(c *gin.Context) {
	jadwalID, err := strconv.Atoi(c.Param("jadwalID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.jadwalService.GetJadwalById(uint(jadwalID))

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		print("msuk sini harusnya")
		return
	}

	c.JSON(200, response)
}

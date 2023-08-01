package controllers

import (
	"strconv"

	"github.com/afdalabdallah/backend-web/models"
	"github.com/afdalabdallah/backend-web/pkg/errs"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

type rumpunController struct {
	rumpunService services.RumpunService
}

type RumpunController interface {
	CreateRumpun(c *gin.Context)
	GetAllRumpun(c *gin.Context)
	DeleteRumpun(c *gin.Context)
	UpdateRumpun(c *gin.Context)
	GetRumpunById(c *gin.Context)
}

func NewRumpunController(rumpunService services.RumpunService) RumpunController {
	return &rumpunController{
		rumpunService: rumpunService,
	}
}

func (p *rumpunController) CreateRumpun(c *gin.Context) {
	var rumpunData models.Rumpun
	if err := c.ShouldBindJSON(&rumpunData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.rumpunService.CreateRumpun(rumpunData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(200, response)

}

func (p *rumpunController) GetAllRumpun(c *gin.Context) {
	response, err := p.rumpunService.GetAllRumpun()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(200, response)
}

func (p *rumpunController) DeleteRumpun(c *gin.Context) {
	rumpunID, err := strconv.Atoi(c.Param("rumpunID"))
	// println(rumpunID)
	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.rumpunService.DeleteRumpun(rumpunID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(200, response)

}

func (p *rumpunController) UpdateRumpun(c *gin.Context) {
	rumpunID, err := strconv.Atoi(c.Param("rumpunID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var rumpunData models.Rumpun
	if err := c.ShouldBindJSON(&rumpunData); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.rumpunService.UpdateRumpun(rumpunID, rumpunData)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(200, response)
}

func (p *rumpunController) GetRumpunById(c *gin.Context) {
	rumpunID, err := strconv.Atoi(c.Param("rumpunID"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())
		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, getErr := p.rumpunService.GetRumpunById(rumpunID)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(200, response)
}

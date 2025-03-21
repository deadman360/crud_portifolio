package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	journey := zap.String("journey", "FindUserByEmail")
	logger.Info("Init FindUserByID route", journey)

	domainResult, err := uc.service.FindUserByID(c.Param("userId"))
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(*domainResult))
}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	journey := zap.String("journey", "FindUserByEmail")
	logger.Info("Init FindUserByEmail route", journey)

	domainResult, err := uc.service.FindUserByEmail(c.Param("userEmail"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(*domainResult))
}

package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserById(c *gin.Context) {
	journey := zap.String("journey", "DeleteUserById")
	logger.Info("Init DeleteUserById route", journey)

	if err := uc.service.DeleteUser(c.Param("userId")); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

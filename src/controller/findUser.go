package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/deadman360/crud_portifolio/src/model"
	"github.com/deadman360/crud_portifolio/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("ID Not Found")
	c.JSON(err.Code, err)
}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	journey := zap.String("journey", "FindUserByEmail")
	logger.Info("Init FindUserByEmail route", journey)

	var UserRequest request.UserRequest
	UserRequest.Email = c.Param("userEmail")

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age)

	domainResult, err := uc.service.FindUserByEmail(domain.GetEmail())
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(*domainResult))
}

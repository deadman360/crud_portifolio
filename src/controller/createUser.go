package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/validation"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/deadman360/crud_portifolio/src/model"
	"github.com/deadman360/crud_portifolio/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))

	var UserRequest request.UserRequest
	if err := c.ShouldBindBodyWithJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully", zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

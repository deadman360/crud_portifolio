package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/validation"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/deadman360/crud_portifolio/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUserById(c *gin.Context) {
	journey := zap.String("journey", "UpdateUserByID")
	logger.Info("Init UpdateUserByID controller route", journey)

	var userRequest request.UserRequest
	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		logger.Error("Error validating Info", err, journey)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.UpdateUser(c.Param("userId"), domain); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

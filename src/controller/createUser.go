package controller

import (
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/validation"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/deadman360/crud_portifolio/src/model"
	"github.com/deadman360/crud_portifolio/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		})
	var UserRequest request.UserRequest
	if err := c.ShouldBindBodyWithJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "CreateUser",
			})
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		})
	c.JSON(http.StatusOK, "")
}

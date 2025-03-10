package controller

import (
	"fmt"
	"net/http"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/validation"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
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
	fmt.Println(UserRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}
	logger.Info("User created sucessfully",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		})
	c.JSON(http.StatusOK, response)
}

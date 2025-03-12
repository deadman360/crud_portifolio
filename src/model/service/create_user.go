package service

import (
	"fmt"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap/zapcore"
)

func (*userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})

	userDomain.EncryptPassword()
	fmt.Println(userDomain)

	return nil
}

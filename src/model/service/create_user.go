package service

import (
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUser")
	logger.Info("Init createUser model", journey)

	userDomain.EncryptPassword()

	//Insert user in database
	userDomainInterfaceRepo, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error creating user", err, journey)
		return nil, err
	}

	return userDomainInterfaceRepo, nil
}

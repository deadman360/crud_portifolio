package service

import (
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByEmailService")
	logger.Info("Starting service for FindUserByEmail", journey)

	userDomainInterfaceRepo, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		restE := rest_err.NewInternalServerError("Error trying to call userRepository", err.Causes)
		return nil, restE
	}

	return &userDomainInterfaceRepo, nil
}

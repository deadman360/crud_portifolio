package service

import (
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(id string) (*model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByIDService")
	logger.Info("Starting service for FindUserByID", journey)

	userDomainInterfaceRepo, err := ud.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return &userDomainInterfaceRepo, nil
}

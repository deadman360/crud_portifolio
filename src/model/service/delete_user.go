package service

import (
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	journey := zap.String("journey", "DeleteUserById")
	logger.Info("Starting service for DeleteUserById", journey)

	if err := ud.userRepository.DeleteUserById(id); err != nil {
		logger.Error("Erro ao chamar o repository", err, journey)
		return err
	}
	return nil
}

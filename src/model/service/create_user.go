package service

import (
	"context"

	"github.com/deadman360/crud_portifolio/src/configuration/database/mongodb"
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/repository"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()
	//Connect to database
	databaseConnection, err := mongodb.NewMongoDbConnection(context.Background())

	if err != nil {
		return rest_err.NewInternalServerError(err.Error(), nil)
	}
	//Insert user in database
	userRepository := repository.NewUserRepository(databaseConnection)
	_, rest_error := userRepository.CreateUser(userDomain)
	if rest_error != nil {
		return rest_error
	}
	return nil
}

package repository

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.uber.org/zap"
)

const (
	MONGO_USER_DB_COLLECTION = "MONGO_DB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUser")
	logger.Info("Init createUser repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), nil)
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), nil)
	}
	userDomain.SetID(result.InsertedID.(string))
	logger.Info("User inserted in database id:"+result.InsertedID.(string), journey)
	return userDomain, nil
}

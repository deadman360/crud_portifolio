package repository

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity/converter"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGO_USER_DB_COLLECTION = "MONGO_USER_DB_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUser")
	logger.Info("Init createUser repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), userEntity)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), nil)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	userEntity.ID = id

	logger.Info("User inserted in database id:"+id, journey)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

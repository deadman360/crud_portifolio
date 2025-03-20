package repository

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity"
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity/converter"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByEmail")
	logger.Info("Init FindUserByEmail repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	//Find in db

	var userEntity entity.UserEntity

	if err := collection.FindOne(context.Background(), bson.M{
		"email": email}, options.FindOne()).Decode(&userEntity); err != nil {
		restE := rest_err.NewBadRequestError("Email not found on our database")
		return nil, restE
	}

	return converter.ConvertEntityToDomain(userEntity), nil
}

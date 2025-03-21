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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByID")
	logger.Info("Init FindUserByID repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	var userEntity entity.UserEntity

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro na conversão do id para objectID", err, journey)
		rest_err := rest_err.NewInternalServerError("Erro na conversão do  ID", nil)
		return nil, rest_err
	}

	if err := collection.FindOne(context.Background(), bson.M{
		"_id": objectID,
	}, options.FindOne()).Decode(&userEntity); err != nil {
		logger.Error("Erro ao dar bind no userEntity", err, journey)
		rest_err := rest_err.NewInternalServerError("ID nao encontrado", nil)
		return nil, rest_err
	}

	return converter.ConvertEntityToDomain(userEntity), nil
}

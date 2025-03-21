package repository

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUserById(id string) *rest_err.RestErr {
	journey := zap.String("journey", "DeleteUserById")
	logger.Info("Init DeleteUserById repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro na conversão de stirng para ObjectID", err, journey)
		rest_err := rest_err.NewInternalServerError("Invalid ID", nil)
		return rest_err
	}
	result, err := collection.DeleteOne(context.Background(), bson.M{
		"_id": objectId,
	}, options.Delete())
	if err != nil {
		logger.Error("Erro ao deletar usuario", err, journey)
		rest_err := rest_err.NewInternalServerError("ID nao encontrado", nil)
		return rest_err
	}

	if result.DeletedCount == 0 {
		logger.Info("DeletedCount = 0", journey)
		rest_err := rest_err.NewInternalServerError("ID nao encontrado", nil)
		return rest_err
	}
	return nil
}

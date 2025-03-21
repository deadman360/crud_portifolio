package repository

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity/converter"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)



func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUser")
	logger.Info("Init createUser repository", journey)

	collection_name := os.Getenv(MONGO_USER_DB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)
	//Verificação pra ver se existe o email na database
	if err := collection.FindOne(context.Background(), bson.M{
		"email": userDomain.GetEmail(),
	}).Err(); err != mongo.ErrNoDocuments {
		return nil, rest_err.NewBadRequestError("Email already in use")
	}

	userEntity := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), userEntity)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), nil)
	}

	userEntity.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("User inserted in database id:"+result.InsertedID.(primitive.ObjectID).Hex(), journey)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

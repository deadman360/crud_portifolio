package repository

import (
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_USER_DB_COLLECTION = "MONGO_USER_DB_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}
type UserRepository interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserById(string) *rest_err.RestErr
}

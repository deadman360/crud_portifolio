package repository

import (
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/model"
	"go.mongodb.org/mongo-driver/mongo"
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
}

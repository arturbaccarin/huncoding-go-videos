package repository

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

type UserRepository interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
}

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

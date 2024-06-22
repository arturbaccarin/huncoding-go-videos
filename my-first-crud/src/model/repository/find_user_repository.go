package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/arturbaccarin/go-my-first-crud/src/model/repository/entity"
	"github.com/arturbaccarin/go-my-first-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("user not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository works",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.Id.Hex()))
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById repository", zap.String("journey", "FindUserById"))

	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("user not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserById repository works",
		zap.String("journey", "FindUserById"),
		zap.String("id", id),
		zap.String("userId", userEntity.Id.Hex()))
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword repository", zap.String("journey", "FindUserByEmailAndPassword"))

	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "user not found with this email and password"
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}

		errorMessage := "error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository works",
		zap.String("journey", "FindUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.Id.Hex()))
	return converter.ConvertEntityToDomain(userEntity), nil
}
